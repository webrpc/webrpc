package ridl

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

var (
	empty = rune(0)
)

type tokenType uint8

const (
	tokenSpace tokenType = iota
	tokenNewLine
	tokenDefine
	tokenDefineValue
	tokenImport
	tokenImportValue
	tokenEnum
	tokenEnumType
	tokenMessage
	tokenService
	tokenField
	tokenMethod
	tokenMethodType
	tokenArgument
	tokenArgumentType
	tokenMeta
	tokenEqual
	tokenWord
	tokenComment
	tokenEOF
)

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

func (lx *lexer) run() {
	log.Printf("running...")
	for state := lexStateStart; state != nil; {
		lx.start = lx.pos
		//log.Printf("state start: %v, lx: %v", stateName(state), lx)
		state = state(lx)
		if lx.pos > lx.start {
			// log.Printf("state: %v, lx: %v", stateName(state), lx)
		}
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

func isAlphanumeric(r rune) bool {
	return !isNewLine(r) && !isSpace(r) && !isEqual(r)
}

func stateName(fn lexState) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
