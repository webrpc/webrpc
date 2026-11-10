package ridl

import (
	"errors"
	"fmt"
	"strings"
)

func expectWord(tok *token, value string) error {
	if tok.tt != tokenWord || tok.val != value {
		return errUnexpectedToken
	}
	return nil
}

func expectNumber(tok *token, value string) error {
	if tok.tt != tokenWord {
		return errUnexpectedToken
	}
	for _, r := range value {
		found := false
		for _, m := range wordNumber {
			if r == m {
				found = true
			}
		}
		if !found {
			return errUnexpectedToken
		}
	}
	return nil
}

func unescapeString(in string) (string, error) {
	size := len(in)
	out := strings.Builder{}

	for i := 0; i < size; i++ {
		c := in[i]

		if c == '\\' {
			i++
			if i < size {
				switch n := in[i]; n {
				case 'n':
					out.WriteString("\n")
				case 't':
					out.WriteString("\t")
				case '"':
					out.WriteByte(n)
				default:
					return "", fmt.Errorf("unexpected character %q after backslash", n)
				}
			} else {
				return "", errors.New("unexpected end after backslash")
			}
		} else {
			// Write the raw byte so multi-byte UTF-8 characters survive.
			out.WriteByte(c)
		}
	}

	return out.String(), nil
}
