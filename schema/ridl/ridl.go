package ridl

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
)

type Tree struct {
}

func Tokenize(input []byte) (*Tree, error) {
	return nil, nil
}
