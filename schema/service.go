package schema

type Service struct {
	Name    VarName  `json:"name"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name    VarName        `json:"name"`
	Inputs  []MethodInput  `json:"inputs"`
	Outputs []MethodOutput `json:"outputs`
}

type MethodInput struct {
	Name VarName  `json:"name"`
	Type DataType `json:"type"`
}

type MethodOutput struct {
	Name VarName  `json:"name"`
	Type DataType `json:"type"`
}
