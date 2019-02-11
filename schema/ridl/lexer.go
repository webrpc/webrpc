package ridl

import (
	"fmt"
	"reflect"
	"runtime"
)

var (
	empty = rune(0)
)

type tokenType uint8

const (
	tokenInvalid tokenType = iota
	tokenSpace
	tokenNewLine
	tokenEqual
	tokenOpenParen
	tokenCloseParen
	tokenPlusSign
	tokenMinusSign
	tokenHash
	tokenWord
	tokenEOF
)

var tokenTypeName = map[tokenType]string{
	tokenInvalid:    "[invalid]",
	tokenSpace:      "[space]",
	tokenNewLine:    "[newline]",
	tokenEqual:      "[equal sign]",
	tokenOpenParen:  "[open parenthesis]",
	tokenCloseParen: "[close parenthesis]",
	tokenPlusSign:   "[plus sign]",
	tokenMinusSign:  "[minus sign]",
	tokenHash:       "[hash sign]",
	tokenWord:       "[word]",
	tokenEOF:        "[EOF]",
}

type token struct {
	tt  tokenType
	val string

	line int
	col  int
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

type lexState func(*lexer) lexState

func (tt tokenType) String() string {
	return tokenTypeName[tt]
}

func (t token) String() string {
	return fmt.Sprintf("%s (line: %d, col: %d): %q", t.tt, t.line, t.col, t.val)
}

func (lx *lexer) run() {
	for state := lexStateStart; state != nil; {
		lx.start = lx.pos
		state = state(lx)
	}
	lx.emit(tokenEOF)
	close(lx.tokens)
}

func Lex(inputString string) *lexer {
	input := []rune(inputString)
	lx := &lexer{
		input:  input,
		length: len(input),
		line:   1,
		tokens: make(chan token, 5),
	}
	go lx.run()
	return lx
}

func (lx *lexer) peek() rune {
	if lx.pos >= lx.length {
		return empty
	}
	return lx.input[lx.pos]
}

func (lx *lexer) next() bool {
	if lx.pos >= lx.length {
		return false
	}
	lx.pos = lx.pos + 1
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

func lexStateSpace(lx *lexer) lexState {
	for lx.next() {
		if !isSpace(lx.peek()) {
			break
		}
	}
	lx.emit(tokenSpace)
	return lexStateStart
}

func lexStateNewLine(lx *lexer) lexState {
	for lx.next() {
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
	lx.next()
	lx.emit(tokenEqual)
	return lexStateStart
}

func lexStateStart(lx *lexer) lexState {
	c := lx.peek()
	switch {
	case isEmpty(c):
		return nil
	case isSpace(c):
		return lexStateSpace
	case isNewLine(c):
		return lexStateNewLine
	case isEqual(c):
		return lexStateEqual
	case isOpenParen(c):
		return lexStateOpenParen
	case isCloseParen(c):
		return lexStateCloseParen
	case isPlusSign(c):
		return lexStatePlusSign
	case isMinusSign(c):
		return lexStateMinusSign
	case isHash(c):
		return lexStateHash
	case isAlphanumeric(c):
		return lexStateWord
	}
	return nil
}

func isEmpty(r rune) bool {
	return r == empty
}

func isSpace(r rune) bool {
	s := string(r)
	if s == " " || s == "\t" || s == "\r" {
		return true
	}
	return false
}

func isNewLine(r rune) bool {
	s := string(r)
	if s == "\n" || s == "\r" {
		return true
	}
	return false
}

func lexStateOpenParen(lx *lexer) lexState {
	lx.next()
	lx.emit(tokenOpenParen)
	return lexStateStart
}

func lexStateCloseParen(lx *lexer) lexState {
	lx.next()
	lx.emit(tokenCloseParen)
	return lexStateStart
}

func lexStatePlusSign(lx *lexer) lexState {
	lx.next()
	lx.emit(tokenPlusSign)
	return lexStateStart
}

func lexStateMinusSign(lx *lexer) lexState {
	lx.next()
	lx.emit(tokenMinusSign)
	return lexStateStart
}

func lexStateHash(lx *lexer) lexState {
	lx.next()
	lx.emit(tokenHash)
	return lexStateStart
}

func lexStateWord(lx *lexer) lexState {
	for lx.next() {
		if !isAlphanumeric(lx.peek()) {
			break
		}
	}
	lx.emit(tokenWord)
	return lexStateStart
}

func isEqual(r rune) bool {
	return string(r) == "="
}

func isPlusSign(r rune) bool {
	return string(r) == "+"
}

func isMinusSign(r rune) bool {
	return string(r) == "-"
}

func isOpenParen(r rune) bool {
	return string(r) == "("
}

func isCloseParen(r rune) bool {
	return string(r) == ")"
}

func isHash(r rune) bool {
	return string(r) == "#"
}

func isAlphanumeric(r rune) bool {
	return !isNewLine(r) && !isSpace(r) && !isEqual(r)
}

func stateName(fn lexState) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
