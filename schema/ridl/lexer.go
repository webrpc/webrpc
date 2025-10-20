package ridl

import (
	"fmt"
)

var (
	empty = rune(0)
)

var (
	wordBeginning = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	wordBreak     = []rune("\x00 \t\r\n[]()<>{}=:¿?¡!,\"")
	wordNumber    = []rune("0123456789")
)

type tokenType uint8

func (tt tokenType) String() string {
	if name := tokenTypeName[tt]; name != "" {
		return name
	}
	return tokenInvalid.String()
}

type token struct {
	tt  tokenType
	val string

	pos  int
	line int
	col  int
}

func (t token) String() string {
	if t.val != "" {
		return t.val
	}
	return t.tt.String()
}

type lexState func(*lexer) lexState

const (
	tokenInvalid           tokenType = iota
	tokenWhitespace                  // " "
	tokenNewLine                     // "\n"
	tokenEqual                       // "="
	tokenOpenParen                   // "("
	tokenCloseParen                  // ")"
	tokenOpenBracket                 // "["
	tokenCloseBracket                // "]"
	tokenOpenAngleBracket            // "<"
	tokenCloseAngleBracket           // ">"
	tokenPlusSign                    // "+"
	tokenMinusSign                   // "-"
	tokenHash                        // "#"
	tokenColon                       // ":"
	tokenComma                       // ","
	tokenBar                         // "|"
	tokenBackslash                   // "\"
	tokenSlash                       // "/"
	tokenQuote                       //  "
	tokenDot                         // "."
	tokenQuestionMark                // "?"
	tokenRocket                      // "=>"
	tokenBang                        // "!"
	tokenWord                        // ..wordCharset..

	tokenExtra // other
	tokenOptionalWhitespace
	tokenComposed

	tokenEOL
	tokenEOF
	tokenAt
)

const tokenDash = tokenMinusSign

var tokenTypeName = map[tokenType]string{
	tokenInvalid:           "[invalid]",
	tokenWhitespace:        "[space]",
	tokenNewLine:           "[newline]",
	tokenEqual:             "[equal sign]",
	tokenOpenParen:         "[open parenthesis]",
	tokenCloseParen:        "[close parenthesis]",
	tokenOpenBracket:       "[open bracket]",
	tokenCloseBracket:      "[close bracket]",
	tokenOpenAngleBracket:  "[open angle bracket]",
	tokenCloseAngleBracket: "[close angle bracket]",
	tokenPlusSign:          "[plus]",
	tokenMinusSign:         "[minus]",
	tokenHash:              "[hash]",
	tokenColon:             "[colon]",
	tokenComma:             "[comma]",
	tokenBar:               "[bar]",
	tokenDot:               "[dot]",
	tokenQuote:             "[quote]",
	tokenBackslash:         "[backslash]",
	tokenSlash:             "[slash]",
	tokenQuestionMark:      "[question mark]",
	tokenRocket:            "[rocket]",
	tokenBang:              "[bang]",
	tokenWord:              "[word]",
	tokenExtra:             "[extra]",
	tokenComposed:          "[composed]",
	tokenEOF:               "[EOF]",
	tokenAt:                "[at]",
}

var tokenTypeValue = map[tokenType][]rune{
	tokenWhitespace:        {' ', '\t', '\r'},
	tokenNewLine:           {'\n'},
	tokenEqual:             {'='},
	tokenOpenParen:         {'('},
	tokenCloseParen:        {')'},
	tokenOpenBracket:       {'['},
	tokenCloseBracket:      {']'},
	tokenOpenAngleBracket:  {'<'},
	tokenCloseAngleBracket: {'>'},
	tokenPlusSign:          {'+'},
	tokenMinusSign:         {'-'},
	tokenHash:              {'#'},
	tokenColon:             {':'},
	tokenQuote:             {'"'},
	tokenBackslash:         {'\\'},
	tokenSlash:             {'/'},
	tokenComma:             {','},
	tokenBar:               {'|'},
	tokenDot:               {'.'},
	tokenQuestionMark:      {'?'},
	tokenBang:              {'!'},
	tokenAt:                {'@'},
}

var (
	isSpace             = isTokenType(tokenWhitespace)
	isNewLine           = isTokenType(tokenNewLine)
	isQuestionMark      = isTokenType(tokenQuestionMark)
	isColon             = isTokenType(tokenColon)
	isHash              = isTokenType(tokenHash)
	isOpenParen         = isTokenType(tokenOpenParen)
	isCloseParen        = isTokenType(tokenCloseParen)
	isOpenBracket       = isTokenType(tokenOpenBracket)
	isCloseBracket      = isTokenType(tokenCloseBracket)
	isOpenAngleBracket  = isTokenType(tokenOpenAngleBracket)
	isCloseAngleBracket = isTokenType(tokenCloseAngleBracket)
	isPlusSign          = isTokenType(tokenPlusSign)
	isMinusSign         = isTokenType(tokenMinusSign)
	isEqual             = isTokenType(tokenEqual)
	isComma             = isTokenType(tokenComma)
	isBar               = isTokenType(tokenBar)
	isQuote             = isTokenType(tokenQuote)
	isBackslash         = isTokenType(tokenBackslash)
	isSlash             = isTokenType(tokenSlash)
	isDot               = isTokenType(tokenDot)
	isBang              = isTokenType(tokenBang)
	isAt                = isTokenType(tokenAt)
)

func isTokenType(tt tokenType) func(r rune) bool {
	return func(r rune) bool {
		for i := range tokenTypeValue[tt] {
			if tokenTypeValue[tt][i] == r {
				return true
			}
		}
		return false
	}
}

func isEmpty(r rune) bool {
	return r == empty
}

func isWordBreak(r rune) bool {
	for i := range wordBreak {
		if r == wordBreak[i] {
			return true
		}
	}
	return false
}

func isWord(r rune) bool {
	for i := range wordBeginning {
		if r == wordBeginning[i] {
			return true
		}
	}
	return false
}

func lexPushTokenState(tt tokenType) lexState {
	return func(lx *lexer) lexState {
		lx.next()
		lx.emit(tt)
		return lexDefaultState
	}
}

func lexStateCloseParen(lx *lexer) lexState {
	return lexPushTokenState(tokenCloseParen)
}

func lexStateOpenParen(lx *lexer) lexState {
	return lexPushTokenState(tokenOpenParen)
}

func lexStateCloseAngleBracket(lx *lexer) lexState {
	return lexPushTokenState(tokenCloseAngleBracket)
}

func lexStateOpenAngleBracket(lx *lexer) lexState {
	return lexPushTokenState(tokenOpenAngleBracket)
}

func lexStateCloseBracket(lx *lexer) lexState {
	return lexPushTokenState(tokenCloseBracket)
}

func lexStateOpenBracket(lx *lexer) lexState {
	return lexPushTokenState(tokenOpenBracket)
}

func lexStateRocket(lx *lexer) lexState {
	return lexPushTokenState(tokenRocket)
}

func lexStateAt(lx *lexer) lexState {
	return lexPushTokenState(tokenAt)
}

func lexStateHash(lx *lexer) lexState {
	return lexPushTokenState(tokenHash)
}

func lexStateComma(lx *lexer) lexState {
	return lexPushTokenState(tokenComma)
}

func lexStateBar(lx *lexer) lexState {
	return lexPushTokenState(tokenBar)
}

func lexStateDot(lx *lexer) lexState {
	return lexPushTokenState(tokenDot)
}

func lexStateExtra(lx *lexer) lexState {
	return lexPushTokenState(tokenExtra)
}

func lexStateColon(lx *lexer) lexState {
	return lexPushTokenState(tokenColon)
}

func lexStateQuestionMark(lx *lexer) lexState {
	return lexPushTokenState(tokenQuestionMark)
}

func lexStatePlusSign(lx *lexer) lexState {
	return lexPushTokenState(tokenPlusSign)
}

func lexStateMinusSign(lx *lexer) lexState {
	return lexPushTokenState(tokenMinusSign)
}

func lexStateQuote(lx *lexer) lexState {
	return lexPushTokenState(tokenQuote)
}

func lexStateSlash(lx *lexer) lexState {
	return lexPushTokenState(tokenSlash)
}

func lexStateBackslash(lx *lexer) lexState {
	return lexPushTokenState(tokenBackslash)
}

func lexStateWord(lx *lexer) lexState {
	for {
		lx.next()
		if isWordBreak(lx.peek()) {
			break
		}
	}

	lx.emit(tokenWord)
	return lexDefaultState
}

func lexStateSpace(lx *lexer) lexState {
	lx.next()

	for isSpace(lx.peek()) {
		lx.next()
	}

	lx.emit(tokenWhitespace)
	return lexDefaultState
}

func lexStateNewLine(lx *lexer) lexState {
	lx.next()
	lx.emit(tokenNewLine)
	lx.col = 0
	return lexDefaultState
}

func lexStateEqual(lx *lexer) lexState {
	lx.next()

	r := lx.peek()

	switch {
	case isCloseAngleBracket(r):
		return lexStateRocket
	}

	lx.emit(tokenEqual)
	return lexDefaultState
}

func lexDefaultState(lx *lexer) lexState {
	r := lx.peek()

	switch {
	case isEmpty(r):
		return nil

	case isQuote(r):
		return lexStateQuote

	case isSlash(r):
		return lexStateSlash

	case isBackslash(r):
		return lexStateBackslash

	case isSpace(r):
		return lexStateSpace

	case isNewLine(r):
		return lexStateNewLine

	case isOpenParen(r):
		return lexStateOpenParen

	case isCloseParen(r):
		return lexStateCloseParen

	case isOpenAngleBracket(r):
		return lexStateOpenAngleBracket

	case isCloseAngleBracket(r):
		return lexStateCloseAngleBracket

	case isOpenBracket(r):
		return lexStateOpenBracket

	case isCloseBracket(r):
		return lexStateCloseBracket

	case isHash(r):
		return lexStateHash

	case isEqual(r):
		return lexStateEqual

	case isPlusSign(r):
		return lexStatePlusSign

	case isMinusSign(r):
		return lexStateMinusSign

	case isColon(r):
		return lexStateColon

	case isQuestionMark(r):
		return lexStateQuestionMark

	case isComma(r):
		return lexStateComma

	case isBar(r):
		return lexStateBar

	case isDot(r):
		return lexStateDot

	case isWord(r):
		return lexStateWord

	case isAt(r):
		return lexStateAt

	default:
		return lexStateExtra
	}
}

type lexer struct {
	input  []rune
	length int

	start int
	pos   int

	line int
	col  int

	tokens chan token
}

func newLexer(in string) *lexer {
	s := []rune(in)
	lx := &lexer{
		input:  s,
		length: len(s),
		tokens: make(chan token),
	}

	go lx.run()
	return lx
}

func (lx *lexer) run() {
	for state := lexDefaultState; state != nil; {
		state = state(lx)
	}

	lx.emit(tokenEOF)
	close(lx.tokens)
}

func (lx *lexer) peek() rune {
	if lx.pos >= lx.length {
		return empty
	}
	return lx.input[lx.pos]
}

func (lx *lexer) next() bool {
	newPos := lx.pos + 1
	if newPos > lx.length {
		return false
	}
	lx.pos = newPos

	if lx.col < 1 {
		lx.line++
	}
	lx.col++

	return true
}

func (lx *lexer) emit(tt tokenType) {
	tok := token{
		tt:   tt,
		val:  lx.val(),
		pos:  lx.pos,
		line: lx.line,
		col:  lx.col,
	}
	lx.start = lx.pos
	lx.tokens <- tok
}

func (lx *lexer) val() string {
	return string(lx.input[lx.start:lx.pos])
}

func (lx *lexer) String() string {
	return fmt.Sprintf("line: %d, start: %d, pos: %d, col: %d, length: %d, value: %q", lx.line, lx.start, lx.pos, lx.col, lx.length, lx.val())
}
