package ridl

import (
	"fmt"
)

var (
	empty = rune(0)
)

var wordCharset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-./")

type tokenType uint8

const (
	tokenInvalid           tokenType = iota
	tokenSpace                       // " "
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
	tokenQuestionMark                // "?"
	tokenRocket                      // "=>"
	tokenWord                        // ..wordCharset..
	tokenExtra                       // other
	tokenComposed
	tokenEOF
)

var tokenTypeName = map[tokenType]string{
	tokenInvalid:           "[invalid]",
	tokenSpace:             "[space]",
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
	tokenQuestionMark:      "[question mark]",
	tokenRocket:            "[rocket]",
	tokenWord:              "[word]",
	tokenExtra:             "[extra]",
	tokenComposed:          "[composed]",
	tokenEOF:               "[EOF]",
}

func (tt tokenType) String() string {
	return tokenTypeName[tt]
}

type token struct {
	tt  tokenType
	val string

	line int
	col  int
}

func (t token) String() string {
	return fmt.Sprintf("%s (line: %d, col: %d): %q", t.tt, t.line, t.col, t.val)
}

type lexer struct {
	input  []rune
	length int

	pos   int
	col   int
	start int
	line  int

	tokens chan token
}

func (lx *lexer) run() {
	for state := lexStateStart; state != nil; {
		lx.start = lx.pos
		state = state(lx)
	}
	lx.emit(tokenEOF)
	close(lx.tokens)
}

func newLexer(inputString string) *lexer {
	input := []rune(inputString)
	lx := &lexer{
		input:  input,
		length: len(input),
		line:   1,
		col:    1,
		tokens: make(chan token, 5),
	}
	go lx.run()
	return lx
}

func (lx *lexer) peek() []rune {
	return lx.peekMore(1)
}

func (lx *lexer) peekMore(chars int) []rune {
	if lx.pos >= lx.length {
		return []rune{empty}
	}
	n := lx.pos + chars
	if lx.length-n <= 0 {
		n = lx.length
	}
	return lx.input[lx.pos:n]
}

func (lx *lexer) next(num int) bool {
	if lx.pos >= lx.length {
		return false
	}
	lx.pos = lx.pos + num
	return true
}

func (lx *lexer) emit(tt tokenType) {
	lx.tokens <- token{
		tt:   tt,
		val:  lx.val(),
		line: lx.line,
		col:  lx.start - lx.col,
	}
}

func (lx *lexer) val() string {
	return string(lx.input[lx.start:lx.pos])
}

func (lx *lexer) String() string {
	return fmt.Sprintf("line: %d, start: %d, pos: %d, col: %d, length: %d, value: %q", lx.line, lx.start, lx.pos, lx.col, lx.length, string(lx.input[lx.start:lx.pos]))
}

type lexState func(*lexer) lexState

func lexStateSpace(lx *lexer) lexState {
	for lx.next(1) {
		if !isSpace(lx.peek()) {
			break
		}
	}
	lx.emit(tokenSpace)
	return lexStateStart
}

func lexStateNewLine(lx *lexer) lexState {
	for lx.next(1) {
		lx.col = lx.pos
		lx.line++
		if !isNewLine(lx.peek()) {
			break
		}
	}
	lx.emit(tokenNewLine)
	return lexStateStart
}

func lexStateEqual(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenEqual)
	return lexStateStart
}

func lexStateStart(lx *lexer) lexState {
	r := lx.peekMore(4)

	switch {
	case isEmpty(r):
		return nil
	case isSpace(r):
		return lexStateSpace
	case isNewLine(r):
		return lexStateNewLine
	case isRocket(r):
		return lexStatePushToken(lx, tokenRocket, 2)
	case isEqual(r):
		return lexStateEqual
	case isOpenParen(r):
		return lexStateOpenParen
	case isCloseParen(r):
		return lexStateCloseParen
	case isPlusSign(r):
		return lexStatePlusSign
	case isMinusSign(r):
		return lexStateMinusSign
	case isHash(r):
		return lexStateHash
	case isColon(r):
		return lexStateColon
	case isQuestionMark(r):
		return lexStateQuestionMark
	case string(r[0]) == "<":
		// TODO: use lexStatePushToken to replace above funcs
		return lexStatePushToken(lx, tokenOpenAngleBracket, 1)
	case string(r[0]) == ">":
		return lexStatePushToken(lx, tokenCloseAngleBracket, 1)
	case string(r[0]) == "[":
		return lexStatePushToken(lx, tokenOpenBracket, 1)
	case string(r[0]) == "]":
		return lexStatePushToken(lx, tokenCloseBracket, 1)
	case string(r[0]) == ",":
		return lexStatePushToken(lx, tokenComma, 1)
	case isWord(r):
		return lexStateWord
	default:
		return lexStateExtra
	}

	panic("reached")
}

func lexStateOpenParen(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenOpenParen)
	return lexStateStart
}

func lexStateCloseParen(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenCloseParen)
	return lexStateStart
}

func lexStateColon(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenColon)
	return lexStateStart
}

func lexStateQuestionMark(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenQuestionMark)
	return lexStateStart
}

func lexStatePlusSign(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenPlusSign)
	return lexStateStart
}

func lexStatePushToken(lx *lexer, tt tokenType, num int) lexState {
	lx.next(num)
	lx.emit(tt)
	return lexStateStart
}

func lexStateMinusSign(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenMinusSign)
	return lexStateStart
}

func lexStateHash(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenHash)
	return lexStateStart
}

func lexStateWord(lx *lexer) lexState {
	for lx.next(1) {
		if !isWord(lx.peek()) {
			break
		}
	}
	lx.emit(tokenWord)
	return lexStateStart
}

func lexStateExtra(lx *lexer) lexState {
	lx.next(1)
	lx.emit(tokenExtra)
	return lexStateStart
}

func isEmpty(r []rune) bool {
	return r[0] == empty
}

func isSpace(r []rune) bool {
	s := string(r[0])
	if s == " " || s == "\t" || s == "\r" {
		return true
	}
	return false
}

func isNewLine(r []rune) bool {
	s := string(r[0:1])
	if s == "\n" || s == "\r" {
		return true
	}
	return false
}

func isEqual(r []rune) bool {
	return string(r[0]) == "="
}

func isPlusSign(r []rune) bool {
	return string(r[0]) == "+"
}

func isMinusSign(r []rune) bool {
	return string(r[0]) == "-"
}

func isOpenParen(r []rune) bool {
	return string(r[0]) == "("
}

func isCloseParen(r []rune) bool {
	return string(r[0]) == ")"
}

func isHash(r []rune) bool {
	return string(r[0]) == "#"
}

func isColon(r []rune) bool {
	return string(r[0]) == ":"
}

func isQuestionMark(r []rune) bool {
	return string(r[0]) == "?"
}

func isRocket(r []rune) bool {
	return len(r) >= 2 && string(r[0:2]) == "=>"
}

func isWord(r []rune) bool {
	for i := range wordCharset {
		if r[0] == wordCharset[i] {
			return true
		}
	}
	return false
}
