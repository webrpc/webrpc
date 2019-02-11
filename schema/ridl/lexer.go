package ridl

import (
	"log"
	"reflect"
	"runtime"
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
	tokenComment
	tokenEOF
)

type lexer struct {
	input  string
	length int

	pos   int
	start int
	line  int

	tokens chan Token
}

type lexState func(*lexer) lexState

func (lx *lexer) run() {
	log.Printf("running...")
	for state := lexStateStart; state != nil; {
		log.Printf("state start: %v", stateName(state))
		state = state(lx)
		log.Printf("state end: %v", stateName(state))
	}
	lx.emit(tokenEOF)
	close(lx.tokens)
}

func (lx *lexer) emit(tokenType) {

}

func Lex(input string) *lexer {
	lx := &lexer{
		input:  input,
		length: len(input),
		tokens: make(chan Token, 5),
	}
	lx.run()
	return lx
}

func (lx *lexer) peek() string {
	if lx.pos >= lx.length {
		return ""
	}
	return lx.input[lx.pos : lx.pos+1]
}

func (lx *lexer) next() string {
	if lx.pos >= lx.length {
		return ""
	}
	s := lx.input[lx.pos : lx.pos+1]
	lx.pos = lx.pos + 1
	return s
}

func lexStateSpace(lx *lexer) lexState {
	for s := lx.next(); isSpace(s); {
		s = lx.next()
	}
	lx.start = lx.pos
	return lexStateStart
}

func lexStateNewLine(lx *lexer) lexState {
	for s := lx.next(); isNewLine(s); {
		s = lx.next()
	}
	lx.start = lx.pos
	return lexStateStart
}

func lexStateStart(lx *lexer) lexState {
	if isSpace(lx.peek()) {
		return lexStateSpace
	}
	if isNewLine(lx.peek()) {
		return lexStateNewLine
	}
	log.Printf("lx: %#v", lx)
	return nil
}

func isSpace(b string) bool {
	if b == " " || b == "\t" || b == "\r" {
		return true
	}
	return false
}

func isNewLine(b string) bool {
	if b == "\n" || b == "\r" {
		return true
	}
	return false
}

func stateName(fn lexState) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
