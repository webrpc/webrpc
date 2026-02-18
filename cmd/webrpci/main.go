package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/pflag"
)

// ---------------------------------------------------------------------------
// webrpc schema types
// ---------------------------------------------------------------------------

type Schema struct {
	WebRPC   string    `json:"webrpc"`
	Name     string    `json:"name"`
	Version  string    `json:"version"`
	Types    []Type    `json:"types"`
	Services []Service `json:"services"`
}

type Type struct {
	Kind     string   `json:"kind"`
	Name     string   `json:"name"`
	TypeStr  string   `json:"type"`
	Fields   []Field  `json:"fields"`
	Comments []string `json:"comments"`
}

type Field struct {
	Name     string            `json:"name"`
	TypeStr  string            `json:"type"`
	Optional bool              `json:"optional"`
	Comments []string          `json:"comments"`
	Meta     []json.RawMessage `json:"meta"`
}

type Service struct {
	Name    string   `json:"name"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name        string                     `json:"name"`
	Succinct    bool                       `json:"succinct"`
	Inputs      []Field                    `json:"inputs"`
	Outputs     []Field                    `json:"outputs"`
	Comments    []string                   `json:"comments"`
	Annotations map[string]json.RawMessage `json:"annotations"`
}

// ---------------------------------------------------------------------------
// History
// ---------------------------------------------------------------------------

type HistoryEntry struct {
	Timestamp   time.Time         `json:"timestamp"`
	Service     string            `json:"service"`
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	RequestBody string            `json:"request_body"`
	Response    string            `json:"response"`
	StatusCode  int               `json:"status_code"`
	Duration    time.Duration     `json:"duration"`
}

func historyPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".webrpci_history")
}

func loadHistory() []HistoryEntry {
	data, err := os.ReadFile(historyPath())
	if err != nil {
		return nil
	}
	var entries []HistoryEntry
	_ = json.Unmarshal(data, &entries)
	return entries
}

func saveHistory(entries []HistoryEntry) {
	if len(entries) > 200 {
		entries = entries[len(entries)-200:]
	}
	data, _ := json.MarshalIndent(entries, "", "  ")
	_ = os.WriteFile(historyPath(), data, 0644)
}

// ---------------------------------------------------------------------------
// JSON syntax highlighting
// ---------------------------------------------------------------------------

var (
	jsonKeyStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#88C0D0"))
	jsonStrStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#A3BE8C"))
	jsonNumStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#B48EAD"))
	jsonBoolStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#D08770"))
	jsonNullStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))
	jsonBraceStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ECEFF4"))
)

func highlightJSON(raw string) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(raw), "", "  "); err != nil {
		return raw
	}
	pretty := buf.String()

	var out strings.Builder
	i := 0

	for i < len(pretty) {
		ch := pretty[i]

		if ch == '"' {
			isKey := isJSONKey(pretty, i)
			end := findStringEnd(pretty, i)
			token := pretty[i : end+1]
			if isKey {
				out.WriteString(jsonKeyStyle.Render(token))
			} else {
				out.WriteString(jsonStrStyle.Render(token))
			}
			i = end + 1
			continue
		}

		if ch == 't' && strings.HasPrefix(pretty[i:], "true") {
			out.WriteString(jsonBoolStyle.Render("true"))
			i += 4
			continue
		}
		if ch == 'f' && strings.HasPrefix(pretty[i:], "false") {
			out.WriteString(jsonBoolStyle.Render("false"))
			i += 5
			continue
		}

		if ch == 'n' && strings.HasPrefix(pretty[i:], "null") {
			out.WriteString(jsonNullStyle.Render("null"))
			i += 4
			continue
		}

		if (ch >= '0' && ch <= '9') || ch == '-' {
			j := i
			for j < len(pretty) && (pretty[j] >= '0' && pretty[j] <= '9' || pretty[j] == '.' || pretty[j] == '-' || pretty[j] == 'e' || pretty[j] == 'E' || pretty[j] == '+') {
				j++
			}
			out.WriteString(jsonNumStyle.Render(pretty[i:j]))
			i = j
			continue
		}

		if ch == '{' || ch == '}' || ch == '[' || ch == ']' {
			out.WriteString(jsonBraceStyle.Render(string(ch)))
			i++
			continue
		}

		out.WriteByte(ch)
		i++
	}

	return out.String()
}

func isJSONKey(s string, pos int) bool {
	end := findStringEnd(s, pos)
	if end+1 >= len(s) {
		return false
	}
	for j := end + 1; j < len(s); j++ {
		if s[j] == ' ' || s[j] == '\t' || s[j] == '\n' || s[j] == '\r' {
			continue
		}
		return s[j] == ':'
	}
	return false
}

func findStringEnd(s string, start int) int {
	esc := false
	for i := start + 1; i < len(s); i++ {
		if esc {
			esc = false
			continue
		}
		if s[i] == '\\' {
			esc = true
			continue
		}
		if s[i] == '"' {
			return i
		}
	}
	return len(s) - 1
}

// ---------------------------------------------------------------------------
// Schema helpers: build request body template from method
// ---------------------------------------------------------------------------

func (s *Schema) typeMap() map[string]Type {
	m := make(map[string]Type, len(s.Types))
	for _, t := range s.Types {
		m[t.Name] = t
	}
	return m
}

func buildTemplate(method Method, types map[string]Type) string {
	if len(method.Inputs) == 0 {
		return "{}"
	}

	obj := make(map[string]interface{})

	if method.Succinct && len(method.Inputs) == 1 {
		typeName := method.Inputs[0].TypeStr
		if t, ok := types[typeName]; ok {
			for _, f := range t.Fields {
				if isHiddenField(f) {
					continue
				}
				obj[f.Name] = defaultForType(f.TypeStr, types, 0)
			}
		}
	} else {
		for _, inp := range method.Inputs {
			obj[inp.Name] = defaultForType(inp.TypeStr, types, 0)
		}
	}

	data, _ := json.MarshalIndent(obj, "", "  ")
	return string(data)
}

func isHiddenField(f Field) bool {
	for _, m := range f.Meta {
		s := string(m)
		if strings.Contains(s, `"json":"-"`) {
			return true
		}
	}
	return false
}

func defaultForType(typeStr string, types map[string]Type, depth int) interface{} {
	if depth > 3 {
		return nil
	}

	if strings.HasPrefix(typeStr, "[]") {
		inner := typeStr[2:]
		return []interface{}{defaultForType(inner, types, depth+1)}
	}

	if strings.HasPrefix(typeStr, "map<") {
		return map[string]interface{}{}
	}

	switch typeStr {
	case "string":
		return ""
	case "bool":
		return false
	case "int", "int32", "int64", "uint8", "uint32", "uint64":
		return 0
	case "float32", "float64":
		return 0.0
	case "bigint":
		return "0"
	case "timestamp":
		return ""
	case "any":
		return nil
	}

	if t, ok := types[typeStr]; ok {
		if t.Kind == "enum" {
			if len(t.Fields) > 0 {
				return t.Fields[0].Name
			}
			return ""
		}
		if t.Kind == "struct" {
			obj := make(map[string]interface{})
			for _, f := range t.Fields {
				if isHiddenField(f) {
					continue
				}
				obj[f.Name] = defaultForType(f.TypeStr, types, depth+1)
			}
			return obj
		}
	}

	return nil
}

// ---------------------------------------------------------------------------
// Service method entry
// ---------------------------------------------------------------------------

type MethodEntry struct {
	ServiceName string
	Method      Method
}

func (m MethodEntry) RPCPath() string {
	return fmt.Sprintf("/rpc/%s/%s", m.ServiceName, m.Method.Name)
}

// ---------------------------------------------------------------------------
// TUI model
// ---------------------------------------------------------------------------

type pane int

const (
	paneMethod pane = iota
	paneRequest
	paneResponse
)

type headerEditState int

const (
	headerEditNone headerEditState = iota
	headerEditKey
	headerEditValue
)

// listLine represents a line in the method list (either a service header or a method).
type listLine struct {
	isHeader    bool
	serviceName string
	methodIdx   int // index into model.methods, only valid when !isHeader
}

type model struct {
	schema   Schema
	types    map[string]Type
	methods  []MethodEntry
	filtered []int // indices into methods

	activePane   pane
	methodCursor int // cursor over filtered methods (not list lines)
	filterMode   bool
	filterText   string

	reqTextarea textarea.Model
	headers     map[string]string

	headerEdit   headerEditState
	headerKeyBuf string
	headerValBuf string

	respViewport viewport.Model
	respContent  string
	respStatus   int
	respDuration time.Duration
	loading      bool

	history       []HistoryEntry
	historyCursor int
	historyMode   bool

	baseURL string

	width  int
	height int
	ready  bool
}

type responseMsg struct {
	body       string
	statusCode int
	duration   time.Duration
	err        error
}

func doRequest(url string, body string, headers map[string]string) tea.Cmd {
	return func() tea.Msg {
		start := time.Now()
		req, err := http.NewRequest("POST", url, strings.NewReader(body))
		if err != nil {
			return responseMsg{err: err}
		}
		req.Header.Set("Content-Type", "application/json")
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		client := &http.Client{Timeout: 30 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return responseMsg{err: err, duration: time.Since(start)}
		}
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		dur := time.Since(start)

		var buf bytes.Buffer
		if json.Indent(&buf, data, "", "  ") == nil {
			return responseMsg{body: buf.String(), statusCode: resp.StatusCode, duration: dur}
		}
		return responseMsg{body: string(data), statusCode: resp.StatusCode, duration: dur}
	}
}

func initialModel(schema Schema, baseURL string) model {
	var methods []MethodEntry
	for _, svc := range schema.Services {
		for _, m := range svc.Methods {
			methods = append(methods, MethodEntry{ServiceName: svc.Name, Method: m})
		}
	}

	ta := textarea.New()
	ta.Placeholder = "Request body (JSON)..."
	ta.CharLimit = 0
	ta.ShowLineNumbers = false

	vp := viewport.New(80, 20)

	types := schema.typeMap()

	filtered := make([]int, len(methods))
	for i := range methods {
		filtered[i] = i
	}

	m := model{
		schema:        schema,
		types:         types,
		methods:       methods,
		filtered:      filtered,
		activePane:    paneMethod,
		reqTextarea:   ta,
		headers:       make(map[string]string),
		respViewport:  vp,
		baseURL:       strings.TrimRight(baseURL, "/"),
		history:       loadHistory(),
		historyCursor: -1,
	}

	if len(methods) > 0 {
		tpl := buildTemplate(methods[0].Method, types)
		m.reqTextarea.SetValue(tpl)
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, tea.EnterAltScreen)
}

func (m model) selectedMethod() *MethodEntry {
	if m.methodCursor < 0 || m.methodCursor >= len(m.filtered) {
		return nil
	}
	return &m.methods[m.filtered[m.methodCursor]]
}

func (m *model) applyFilter() {
	m.filtered = m.filtered[:0]
	lower := strings.ToLower(m.filterText)
	for i, me := range m.methods {
		name := strings.ToLower(me.ServiceName + "/" + me.Method.Name)
		if strings.Contains(name, lower) {
			m.filtered = append(m.filtered, i)
		}
	}
	if m.methodCursor >= len(m.filtered) {
		m.methodCursor = 0
	}
}

func (m *model) loadMethodTemplate() {
	me := m.selectedMethod()
	if me == nil {
		return
	}
	tpl := buildTemplate(me.Method, m.types)
	m.reqTextarea.SetValue(tpl)
	m.historyCursor = -1
	m.historyMode = false
}

func (m *model) methodHistory() []HistoryEntry {
	me := m.selectedMethod()
	if me == nil {
		return nil
	}
	var result []HistoryEntry
	for _, h := range m.history {
		if h.Service == me.ServiceName && h.Method == me.Method.Name {
			result = append(result, h)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Timestamp.After(result[j].Timestamp)
	})
	return result
}

// buildListLines builds the display lines for the method pane, interleaving
// service headers and methods. It returns the lines and a mapping from
// cursor position to list line index so we can scroll properly.
func (m model) buildListLines(w int) (lines []string, cursorLine int) {
	type group struct {
		name    string
		indices []int
	}
	var groups []group
	groupMap := make(map[string]int)

	for _, idx := range m.filtered {
		me := m.methods[idx]
		gi, ok := groupMap[me.ServiceName]
		if !ok {
			gi = len(groups)
			groupMap[me.ServiceName] = gi
			groups = append(groups, group{name: me.ServiceName})
		}
		groups[gi].indices = append(groups[gi].indices, idx)
	}

	cursorLine = 0
	methodIdx := 0
	for _, g := range groups {
		lines = append(lines, svcHeaderStyle.Render(truncPad(g.name, w)))
		for _, idx := range g.indices {
			me := m.methods[idx]
			name := me.Method.Name

			// Show input count hint
			nInputs := countInputs(me.Method, m.types)
			suffix := ""
			if nInputs > 0 {
				suffix = dimStyle.Render(fmt.Sprintf(" (%d)", nInputs))
			}

			prefix := "  "
			if methodIdx == m.methodCursor {
				cursorLine = len(lines)
				line := selectedStyle.Render(truncPad(prefix+name, w))
				if suffix != "" {
					// Append hint after selected highlight
					line += suffix
				}
				lines = append(lines, line)
			} else {
				deprecated := false
				if _, ok := me.Method.Annotations["deprecated"]; ok {
					deprecated = true
				}
				display := prefix + name
				if deprecated {
					lines = append(lines, dimStyle.Render(truncPad(display, w))+suffix)
				} else {
					lines = append(lines, truncPad(display, w)+suffix)
				}
			}
			methodIdx++
		}
	}
	return
}

func countInputs(method Method, types map[string]Type) int {
	if len(method.Inputs) == 0 {
		return 0
	}
	if method.Succinct && len(method.Inputs) == 1 {
		typeName := method.Inputs[0].TypeStr
		if t, ok := types[typeName]; ok {
			count := 0
			for _, f := range t.Fields {
				if !isHiddenField(f) {
					count++
				}
			}
			return count
		}
	}
	return len(method.Inputs)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		m.recalcLayout()
		return m, nil

	case responseMsg:
		m.loading = false
		if msg.err != nil {
			m.respContent = fmt.Sprintf("Error: %s", msg.err)
			m.respStatus = 0
		} else {
			m.respContent = msg.body
			m.respStatus = msg.statusCode
		}
		m.respDuration = msg.duration
		highlighted := highlightJSON(m.respContent)
		m.respViewport.SetContent(highlighted)
		m.respViewport.GotoTop()

		me := m.selectedMethod()
		if me != nil {
			entry := HistoryEntry{
				Timestamp:   time.Now(),
				Service:     me.ServiceName,
				Method:      me.Method.Name,
				URL:         m.baseURL + me.RPCPath(),
				Headers:     copyMap(m.headers),
				RequestBody: m.reqTextarea.Value(),
				Response:    m.respContent,
				StatusCode:  m.respStatus,
				Duration:    m.respDuration,
			}
			m.history = append(m.history, entry)
			saveHistory(m.history)
		}
		return m, nil

	case tea.KeyMsg:
		key := msg.String()

		// Header editing mode intercepts all keys
		if m.headerEdit != headerEditNone {
			return m.updateHeaderEdit(msg)
		}

		// Global keys
		switch key {
		case "ctrl+c":
			return m, tea.Quit
		case "f2":
			// Toggle header editing
			m.headerEdit = headerEditKey
			m.headerKeyBuf = ""
			m.headerValBuf = ""
			return m, nil
		}

		// Filter mode
		if m.filterMode {
			switch key {
			case "esc":
				m.filterMode = false
				m.filterText = ""
				m.applyFilter()
				return m, nil
			case "enter":
				m.filterMode = false
				return m, nil
			case "backspace":
				if len(m.filterText) > 0 {
					m.filterText = m.filterText[:len(m.filterText)-1]
					m.applyFilter()
				}
				return m, nil
			default:
				if len(key) == 1 {
					m.filterText += key
					m.applyFilter()
				}
				return m, nil
			}
		}

		// Pane-specific handling
		switch m.activePane {
		case paneMethod:
			switch key {
			case "q":
				return m, tea.Quit
			case "tab":
				m.activePane = paneRequest
				m.reqTextarea.Focus()
				return m, nil
			case "up", "k":
				if m.methodCursor > 0 {
					m.methodCursor--
					m.loadMethodTemplate()
				}
				return m, nil
			case "down", "j":
				if m.methodCursor < len(m.filtered)-1 {
					m.methodCursor++
					m.loadMethodTemplate()
				}
				return m, nil
			case "home", "g":
				m.methodCursor = 0
				m.loadMethodTemplate()
				return m, nil
			case "end", "G":
				if len(m.filtered) > 0 {
					m.methodCursor = len(m.filtered) - 1
					m.loadMethodTemplate()
				}
				return m, nil
			case "/":
				m.filterMode = true
				m.filterText = ""
				return m, nil
			case "enter":
				m.activePane = paneRequest
				m.reqTextarea.Focus()
				return m, nil
			case "ctrl+x":
				return m.executeRequest()
			}

		case paneRequest:
			switch key {
			case "esc":
				m.activePane = paneMethod
				m.reqTextarea.Blur()
				m.historyMode = false
				return m, nil
			case "tab":
				m.activePane = paneResponse
				m.reqTextarea.Blur()
				return m, nil
			case "ctrl+x":
				return m.executeRequest()
			case "ctrl+r":
				// Reset to template
				m.loadMethodTemplate()
				return m, nil
			case "ctrl+p":
				hist := m.methodHistory()
				if len(hist) > 0 {
					m.historyMode = true
					m.historyCursor++
					if m.historyCursor >= len(hist) {
						m.historyCursor = len(hist) - 1
					}
					m.reqTextarea.SetValue(hist[m.historyCursor].RequestBody)
				}
				return m, nil
			case "ctrl+n":
				hist := m.methodHistory()
				if m.historyMode && len(hist) > 0 {
					m.historyCursor--
					if m.historyCursor < 0 {
						m.historyCursor = -1
						m.historyMode = false
						m.loadMethodTemplate()
					} else {
						m.reqTextarea.SetValue(hist[m.historyCursor].RequestBody)
					}
				}
				return m, nil
			default:
				var cmd tea.Cmd
				m.reqTextarea, cmd = m.reqTextarea.Update(msg)
				cmds = append(cmds, cmd)
				return m, tea.Batch(cmds...)
			}

		case paneResponse:
			switch key {
			case "esc":
				m.activePane = paneMethod
				return m, nil
			case "tab":
				m.activePane = paneMethod
				return m, nil
			case "shift+tab":
				m.activePane = paneRequest
				m.reqTextarea.Focus()
				return m, nil
			case "q":
				return m, tea.Quit
			case "ctrl+x":
				return m.executeRequest()
			default:
				var cmd tea.Cmd
				m.respViewport, cmd = m.respViewport.Update(msg)
				cmds = append(cmds, cmd)
				return m, tea.Batch(cmds...)
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) updateHeaderEdit(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	switch key {
	case "esc":
		m.headerEdit = headerEditNone
		return m, nil
	case "enter":
		if m.headerEdit == headerEditKey {
			if m.headerKeyBuf == "" {
				m.headerEdit = headerEditNone
				return m, nil
			}
			m.headerEdit = headerEditValue
			return m, nil
		}
		if m.headerValBuf == "" {
			delete(m.headers, m.headerKeyBuf)
		} else {
			m.headers[m.headerKeyBuf] = m.headerValBuf
		}
		m.headerEdit = headerEditNone
		return m, nil
	case "backspace":
		if m.headerEdit == headerEditKey && len(m.headerKeyBuf) > 0 {
			m.headerKeyBuf = m.headerKeyBuf[:len(m.headerKeyBuf)-1]
		} else if m.headerEdit == headerEditValue && len(m.headerValBuf) > 0 {
			m.headerValBuf = m.headerValBuf[:len(m.headerValBuf)-1]
		}
		return m, nil
	default:
		if len(key) == 1 || key == " " {
			if m.headerEdit == headerEditKey {
				m.headerKeyBuf += key
			} else {
				m.headerValBuf += key
			}
		}
		return m, nil
	}
}

func (m *model) executeRequest() (model, tea.Cmd) {
	me := m.selectedMethod()
	if me == nil {
		return *m, nil
	}
	url := m.baseURL + me.RPCPath()
	body := m.reqTextarea.Value()
	m.loading = true
	m.respContent = "Loading..."
	m.respViewport.SetContent(loadingStyle.Render("Loading..."))
	m.activePane = paneResponse
	m.reqTextarea.Blur()
	return *m, doRequest(url, body, m.headers)
}

func (m *model) recalcLayout() {
	if !m.ready {
		return
	}
	contentHeight := m.height - 5
	if contentHeight < 5 {
		contentHeight = 5
	}

	_, midW, rightW := m.paneWidths()

	taHeight := contentHeight - 4
	if taHeight < 3 {
		taHeight = 3
	}

	m.reqTextarea.SetWidth(midW - 4)
	m.reqTextarea.SetHeight(taHeight)

	m.respViewport.Width = rightW - 4
	m.respViewport.Height = contentHeight
}

func (m model) paneWidths() (int, int, int) {
	w := m.width - 4
	if w < 30 {
		w = 30
	}
	left := w * 22 / 100
	if left < 20 {
		left = 20
	}
	if left > 35 {
		left = 35
	}
	mid := w * 38 / 100
	if mid < 25 {
		mid = 25
	}
	right := w - left - mid
	if right < 20 {
		right = 20
	}
	return left, mid, right
}

// ---------------------------------------------------------------------------
// View
// ---------------------------------------------------------------------------

var (
	activeBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#88C0D0"))
	inactiveBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#4C566A"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#88C0D0"))

	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#2E3440")).
			Background(lipgloss.Color("#88C0D0"))

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#616E88"))

	svcHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#EBCB8B")).
			PaddingLeft(0)

	statusOKStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A3BE8C"))
	statusErrStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#BF616A"))

	headerPromptStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#D08770"))

	loadingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EBCB8B")).
			Bold(true)

	historyActiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#B48EAD"))

	urlStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5E81AC"))
)

func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	leftW, midW, rightW := m.paneWidths()
	contentHeight := m.height - 5
	if contentHeight < 5 {
		contentHeight = 5
	}

	leftContent := m.renderMethodPane(leftW-2, contentHeight)
	midContent := m.renderRequestPane(midW-2, contentHeight)
	rightContent := m.renderResponsePane(rightW-2, contentHeight)

	leftBorder := inactiveBorderStyle
	midBorder := inactiveBorderStyle
	rightBorder := inactiveBorderStyle

	switch m.activePane {
	case paneMethod:
		leftBorder = activeBorderStyle
	case paneRequest:
		midBorder = activeBorderStyle
	case paneResponse:
		rightBorder = activeBorderStyle
	}

	left := leftBorder.Width(leftW).Height(contentHeight).Render(leftContent)
	mid := midBorder.Width(midW).Height(contentHeight).Render(midContent)
	right := rightBorder.Width(rightW).Height(contentHeight).Render(rightContent)

	mainView := lipgloss.JoinHorizontal(lipgloss.Top, left, mid, right)

	header := m.renderHeader()
	footer := m.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, mainView, footer)
}

func (m model) renderHeader() string {
	me := m.selectedMethod()
	var methodInfo string
	if me != nil {
		path := me.RPCPath()
		methodInfo = "  " + urlStyle.Render(m.baseURL+path)
		if len(me.Method.Comments) > 0 {
			methodInfo += "  " + dimStyle.Render(me.Method.Comments[0])
		}
	}
	return titleStyle.Render("webrpci") +
		dimStyle.Render(" "+m.schema.Name+" "+m.schema.Version) +
		methodInfo
}

func (m model) renderFooter() string {
	if m.headerEdit == headerEditKey {
		return headerPromptStyle.Render("Header name (empty=cancel): ") + m.headerKeyBuf + "\u2588"
	}
	if m.headerEdit == headerEditValue {
		return headerPromptStyle.Render(m.headerKeyBuf+": ") + m.headerValBuf + "\u2588"
	}
	if m.filterMode {
		return headerPromptStyle.Render("/ ") + m.filterText + "\u2588"
	}

	parts := []string{
		dimStyle.Render("tab") + ":pane",
		dimStyle.Render("^x") + ":run",
		dimStyle.Render("F2") + ":header",
		dimStyle.Render("/") + ":filter",
		dimStyle.Render("^p/^n") + ":history",
		dimStyle.Render("^r") + ":reset",
		dimStyle.Render("q") + ":quit",
	}

	var extra string
	if m.respStatus > 0 {
		s := fmt.Sprintf("  HTTP %d", m.respStatus)
		if m.respStatus >= 200 && m.respStatus < 300 {
			extra = statusOKStyle.Render(s)
		} else {
			extra = statusErrStyle.Render(s)
		}
		extra += dimStyle.Render(fmt.Sprintf(" %s", m.respDuration.Round(time.Millisecond)))
	}
	if m.loading {
		extra = loadingStyle.Render("  requesting...")
	}

	return strings.Join(parts, dimStyle.Render(" | ")) + extra
}

func (m model) renderMethodPane(w, h int) string {
	lines, cursorLine := m.buildListLines(w)

	// Add filter indicator at top if active
	var prefix []string
	if m.filterMode || m.filterText != "" {
		prefix = append(prefix, dimStyle.Render("filter: ")+m.filterText)
	}

	all := append(prefix, lines...)

	// Adjust cursorLine for prefix offset
	cursorLine += len(prefix)

	// Pad to height
	for len(all) < h {
		all = append(all, "")
	}

	// Scroll to keep cursor visible
	if len(all) > h {
		start := 0
		if cursorLine > h-3 {
			start = cursorLine - h + 3
		}
		end := start + h
		if end > len(all) {
			end = len(all)
			start = end - h
			if start < 0 {
				start = 0
			}
		}
		all = all[start:end]
	}

	return strings.Join(all, "\n")
}

func (m model) renderRequestPane(w, h int) string {
	var lines []string
	headerLines := 0

	// Show custom headers
	if len(m.headers) > 0 {
		lines = append(lines, dimStyle.Render("Headers:"))
		headerLines++
		keys := make([]string, 0, len(m.headers))
		for k := range m.headers {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			v := m.headers[k]
			hl := fmt.Sprintf("  %s: %s", k, v)
			if len(hl) > w {
				hl = hl[:w-3] + "..."
			}
			lines = append(lines, dimStyle.Render(hl))
			headerLines++
		}
		lines = append(lines, "")
		headerLines++
	}

	// Show history hint
	hist := m.methodHistory()
	if m.historyMode && m.historyCursor >= 0 && m.historyCursor < len(hist) {
		hEntry := hist[m.historyCursor]
		ts := hEntry.Timestamp.Format("15:04:05")
		status := ""
		if hEntry.StatusCode > 0 {
			if hEntry.StatusCode >= 200 && hEntry.StatusCode < 300 {
				status = statusOKStyle.Render(fmt.Sprintf(" %d", hEntry.StatusCode))
			} else {
				status = statusErrStyle.Render(fmt.Sprintf(" %d", hEntry.StatusCode))
			}
		}
		label := historyActiveStyle.Render(fmt.Sprintf("history [%d/%d] %s", m.historyCursor+1, len(hist), ts)) + status
		lines = append(lines, label)
		headerLines++
	} else if len(hist) > 0 {
		lines = append(lines, dimStyle.Render(fmt.Sprintf("(%d in history, ^p to browse)", len(hist))))
		headerLines++
	}

	// Body label
	lines = append(lines, dimStyle.Render("Body:"))
	headerLines++

	remainingHeight := h - headerLines - 1
	if remainingHeight < 3 {
		remainingHeight = 3
	}
	m.reqTextarea.SetHeight(remainingHeight)
	m.reqTextarea.SetWidth(w)

	lines = append(lines, m.reqTextarea.View())

	return strings.Join(lines, "\n")
}

func (m model) renderResponsePane(w, h int) string {
	if m.respContent == "" {
		help := []string{
			dimStyle.Render("No response yet."),
			"",
			dimStyle.Render("Select a method and press"),
			dimStyle.Render("ctrl+x to execute."),
			"",
			dimStyle.Render("Navigation:"),
			dimStyle.Render("  tab     - switch panes"),
			dimStyle.Render("  j/k     - methods up/down"),
			dimStyle.Render("  /       - filter methods"),
			dimStyle.Render("  enter   - edit request"),
			dimStyle.Render("  ctrl+x  - execute request"),
			dimStyle.Render("  ctrl+p  - prev history"),
			dimStyle.Render("  ctrl+n  - next history"),
			dimStyle.Render("  ctrl+r  - reset body"),
			dimStyle.Render("  F2      - add header"),
			dimStyle.Render("  q       - quit"),
		}
		return strings.Join(help, "\n")
	}

	m.respViewport.Width = w
	m.respViewport.Height = h
	return m.respViewport.View()
}

func truncPad(s string, max int) string {
	if len(s) > max {
		if max > 3 {
			return s[:max-3] + "..."
		}
		return s[:max]
	}
	return s
}

func copyMap(m map[string]string) map[string]string {
	c := make(map[string]string, len(m))
	for k, v := range m {
		c[k] = v
	}
	return c
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	schemaFile := pflag.StringP("schema", "s", "", "Path to webrpc .gen.json schema file")
	baseURL := pflag.StringP("url", "u", "http://localhost:8080", "Base URL of the webrpc server")
	pflag.Parse()

	if *schemaFile == "" {
		fmt.Fprintln(os.Stderr, "Usage: webrpci --schema <file.gen.json> --url <base-url>")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Example:")
		fmt.Fprintln(os.Stderr, "  webrpci -s ./proto/api.gen.json -u http://localhost:4424")
		os.Exit(1)
	}

	data, err := os.ReadFile(*schemaFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading schema: %s\n", err)
		os.Exit(1)
	}

	var schema Schema
	if err := json.Unmarshal(data, &schema); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing schema: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Loading %s %s (%d services, %d types)\n",
		schema.Name, schema.Version, len(schema.Services), len(schema.Types))

	m := initialModel(schema, *baseURL)

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
