package ridl

import (
	"errors"
	"fmt"
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
	out := ""

	for i := 0; i < size; i++ {
		c := in[i]

		if c == '\\' {
			i++
			if i < size {
				switch n := in[i]; n {
				case 'n':
					out = out + "\n"
				case 't':
					out = out + "\t"
				case '"':
					out = out + string(n)
				default:
					return "", fmt.Errorf("unexpected character %q after backslash", c)
				}
			} else {
				return "", errors.New("unexpected end after backslash")
			}
		} else {
			out = out + string(c)
		}
	}

	return out, nil
}
