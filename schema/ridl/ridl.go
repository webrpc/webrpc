package ridl

type Tree struct {
}

type Token struct {
}

func Tokenize(input []byte) (*Tree, error) {
	_ = Lex(string(input))
	return nil, nil
}
