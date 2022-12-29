package ridl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerColumnAndLine(t *testing.T) {
	type lc struct {
		line int
		col  int
	}

	inputs := []struct {
		in  string
		out []lc
	}{
		{
			"",
			[]lc{
				{0, 0}, // ""
				{0, 0}, // EOF
			},
		},
		{
			"A",
			[]lc{
				{1, 1}, // A
				{1, 1}, // EOF
			},
		},
		{
			" ",
			[]lc{
				{1, 1}, // " "
				{1, 1}, // EOF
			},
		},
		{
			" ABC",
			[]lc{
				{1, 1}, // " "
				{1, 4}, // ABC
				{1, 4}, // EOF
			},
		},
		{
			" ABC\nZ\nDEF ",
			[]lc{
				{1, 1}, // " "
				{1, 4}, // ABC
				{1, 5}, // \n
				{2, 1}, // Z
				{2, 2}, // \n
				{3, 3}, // DEF
				{3, 4}, // " "
				{3, 4}, // EOF
			},
		},
		{
			" ABC\n\n\nZ\nDEF ",
			[]lc{
				{1, 1}, // " "
				{1, 4}, // ABC
				{1, 5}, // \n
				{2, 1}, // \n
				{3, 1}, // \n
				{4, 1}, // Z
				{4, 2}, // \n
				{5, 3}, // DEF
				{5, 4}, // " "
				{5, 4}, // EOF
			},
		},
		{
			"=\nA\nC\nD EFG HIJK L",
			[]lc{
				{1, 1},  // =
				{1, 2},  // \n
				{2, 1},  // A
				{2, 2},  // \n
				{3, 1},  // C
				{3, 2},  // \n
				{4, 1},  // D
				{4, 2},  // " "
				{4, 5},  // EFG
				{4, 6},  // " "
				{4, 10}, // HIJK
				{4, 11}, // " "
				{4, 12}, // L
				{4, 12}, // EOF
			},
		},
		{
			"=",
			[]lc{
				{1, 1}, // =
				{1, 1}, // EOF
			},
		},
		{
			"=>",
			[]lc{
				{1, 2}, // =>
				{1, 2}, //
			},
		},
		{
			"=> =>",
			[]lc{
				{1, 2}, // =>
				{1, 3}, // " "
				{1, 5}, // =>
				{1, 5}, // EOF
			},
		},
		{
			"=> =<",
			[]lc{
				{1, 2}, // =>
				{1, 3}, // " "
				{1, 4}, // =
				{1, 5}, // <
				{1, 5}, // EOF
			},
		},
		{
			"=>    =<",
			[]lc{
				{1, 2}, // =>
				{1, 6}, // "    "
				{1, 7}, // =
				{1, 8}, // <
				{1, 8}, // EOF
			},
		},
		{
			"=>   ... =>   =<",
			[]lc{
				{1, 2},  // =>
				{1, 5},  // "    "
				{1, 6},  // .
				{1, 7},  // .
				{1, 8},  // .
				{1, 9},  // " "
				{1, 11}, // =>
				{1, 14}, // "    "
				{1, 15}, // =
				{1, 16}, // <
				{1, 16}, // EOF
			},
		},
		{
			"=>=>==>",
			[]lc{
				{1, 2}, // =>
				{1, 4}, // =>
				{1, 5}, // =
				{1, 7}, // =>
				{1, 7}, // EOF
			},
		},
		{
			`""`,
			[]lc{
				{1, 1},
				{1, 2},
				{1, 2},
			},
		},
		{
			`"abc\"`,
			[]lc{
				{1, 1},
				{1, 5},
				{1, 6},
				{1, 6},
			},
		},
	}

	for _, input := range inputs {
		lx := newLexer(input.in)
		assert.NotNil(t, lx)

		i := 0
		for tok := range lx.tokens {
			if !assert.True(t, i < len(input.out)) {
				t.Fatalf("got more tokens than expected: %#v", tok)
			}
			assert.Equal(t, input.out[i].line, tok.line, "line mismatch (input: %q)", input.in)
			assert.Equal(t, input.out[i].col, tok.col, "column mismatch (input: %q)", input.in)
			i++
		}
	}
}

func TestLexerEmmitter(t *testing.T) {
	inputs := []struct {
		in  string
		out []tokenType
	}{
		{
			"",
			[]tokenType{
				tokenEOF,
			},
		},
		{
			"=",
			[]tokenType{
				tokenEqual,
				tokenEOF,
			},
		},
		{
			"===",
			[]tokenType{
				tokenEqual,
				tokenEqual,
				tokenEqual,
				tokenEOF,
			},
		},
		{
			"=ABC=",
			[]tokenType{
				tokenEqual,
				tokenWord,
				tokenEqual,
				tokenEOF,
			},
		},
		{
			"...===>=>",
			[]tokenType{
				tokenDot,
				tokenDot,
				tokenDot,
				tokenEqual,
				tokenEqual,
				tokenRocket,
				tokenRocket,
				tokenEOF,
			},
		},
		{
			`"abc\""`,
			[]tokenType{
				tokenQuote,
				tokenWord,
				tokenQuote,
				tokenQuote,
				tokenEOF,
			},
		},
	}

	for _, input := range inputs {
		lx := newLexer(input.in)
		assert.NotNil(t, lx)

		i := 0
		for tok := range lx.tokens {
			if !assert.True(t, i < len(input.out)) {
				t.Fatalf("got more tokens than expected: %#v", tok)
			}
			assert.Equal(t, input.out[i], tok.tt)
			i++
		}
	}
}

func TestLexerStringTokens(t *testing.T) {
	inputs := []struct {
		in  string
		out []string
	}{
		{
			"",
			[]string{},
		},
		{
			`""\"`,
			[]string{
				`"`,
				`"`,
				`\`,
				`"`,
			},
		},
		{
			"=ABC",
			[]string{
				"=",
				"ABC",
			},
		},
		{
			"=> =>",
			[]string{
				"=>",
				" ",
				"=>",
			},
		},
		{
			"=> ==>",
			[]string{
				"=>",
				" ",
				"=",
				"=>",
			},
		},
		{
			"=>=== >",
			[]string{
				"=>",
				"=",
				"=",
				"=",
				" ",
				">",
			},
		},
		{
			"=ABC<>\nDEF=(=)==>=>\n=\n\n\n=>=>==>...ABC\n\n",
			[]string{
				"=",
				"ABC",
				"<",
				">",
				"\n",
				"DEF",
				"=",
				"(",
				"=",
				")",
				"=",
				"=>",
				"=>",
				"\n",
				"=",
				"\n",
				"\n",
				"\n",
				"=>",
				"=>",
				"=",
				"=>",
				".",
				".",
				".",
				"ABC",
				"\n",
				"\n",
			},
		},
		{
			" \n\n\t\twebrpc    =v1 + foo = bar",
			[]string{
				" ",
				"\n",
				"\n",
				"\t\t",
				"webrpc",
				"    ",
				"=",
				"v1",
				" ",
				"+",
				" ",
				"foo",
				" ",
				"=",
				" ",
				"bar",
			},
		},
	}

	for _, input := range inputs {
		lx := newLexer(input.in)
		assert.NotNil(t, lx)

		i := 0
		for tok := range lx.tokens {
			if tok.tt != tokenEOF {
				if !assert.True(t, i < len(input.out)) {
					t.Fatalf("got more tokens than expected: %#v", tok)
				}
				assert.Equal(t, input.out[i], tok.val)
			}
			i++
		}
	}
}

func TestLexerSimpleTokens(t *testing.T) {
	inputs := []struct {
		in  string
		out []tokenType
	}{
		{
			"",
			[]tokenType{},
		},
		{
			" ",
			[]tokenType{tokenWhitespace},
		},
		{
			"AAAA",
			[]tokenType{
				tokenWord,
			},
		},
		{
			" \tAAAA       ",
			[]tokenType{
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
			},
		},
		{
			" -AAAA + - +-",
			[]tokenType{
				tokenWhitespace,
				tokenMinusSign,
				tokenWord,
				tokenWhitespace,
				tokenPlusSign,
				tokenWhitespace,
				tokenMinusSign,
				tokenWhitespace,
				tokenPlusSign,
				tokenMinusSign,
			},
		},
		{
			"- -A - --( AA?:AA AA:? BB AA-A-A-A-A-",
			[]tokenType{
				tokenMinusSign,
				tokenWhitespace,
				tokenMinusSign,
				tokenWord,
				tokenWhitespace,
				tokenMinusSign,
				tokenWhitespace,
				tokenMinusSign,
				tokenMinusSign,
				tokenOpenParen,
				tokenWhitespace,
				tokenWord,
				tokenQuestionMark,
				tokenColon,
				tokenWord,
				tokenWhitespace,
				tokenWord,
				tokenColon,
				tokenQuestionMark,
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenWord,
			},
		},
		{
			"AAA BBBB CCC =",
			[]tokenType{
				tokenWord,
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenEqual,
			},
		},
		{
			"   AAA    \t\t         BBBB CCC =\t\t=>",
			[]tokenType{
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenEqual,
				tokenWhitespace,
				tokenRocket,
			},
		},
		{
			"  \t\t\t\t AAA    \t\t         BBBB CCC          \t\t\t\t =",
			[]tokenType{
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenWord,
				tokenWhitespace,
				tokenEqual,
			},
		},
		{
			"A\n\tB",
			[]tokenType{
				tokenWord,
				tokenNewLine,
				tokenWhitespace,
				tokenWord,
			},
		},
		{
			`""" AA=\`,
			[]tokenType{
				tokenQuote,
				tokenQuote,
				tokenQuote,
				tokenWhitespace,
				tokenWord,
				tokenEqual,
				tokenBackslash,
			},
		},
	}

	for _, input := range inputs {
		tokens, err := tokenize([]byte(input.in))
		assert.NoError(t, err)

		assert.Equal(t, len(input.out), len(tokens))
		for i, token := range tokens {
			if assert.NotZero(t, input.out[i]) {
				assert.Equal(t, input.out[i], token.tt, "expecting: %s, got: %s, near: %d (input: %q)", input.out[i], token.tt, i, input.in)
			}
		}
	}
}

func TestLexerRIDLTokens(t *testing.T) {
	inputs := []struct {
		in  string
		out []string
	}{
		{
			"\n\n\nwebrpc",
			[]string{
				"\n",
				"\n",
				"\n",
				"webrpc",
			},
		},
		{
			"\n\n\nwebrpc\n=v1",
			[]string{
				"\n",
				"\n",
				"\n",
				"webrpc",
				"\n",
				"=",
				"v1",
			},
		},
		{
			"\n\n\nwebrpc\n=v1\n\t\t\t+   foo = bar - baz = 56 # a    comment\n\nversion = v0.0.0.1",
			[]string{
				"\n",
				"\n",
				"\n",
				"webrpc",
				"\n",
				"=",
				"v1",
				"\n",
				"\t\t\t",
				"+",
				"   ",
				"foo",
				" ",
				"=",
				" ",
				"bar",
				" ",
				"-",
				" ",
				"baz",
				" ",
				"=",
				" ",
				"56",
				" ",
				"#",
				" ",
				"a",
				"    ",
				"comment",
				"\n",
				"\n",
				"version",
				" ",
				"=",
				" ",
				"v0.0.0.1",
			},
		},
	}

	for _, input := range inputs {
		tokens, err := tokenize([]byte(input.in))
		assert.NoError(t, err)

		for i, tok := range tokens {
			assert.Equal(t, input.out[i], tok.val)
		}
	}
}
