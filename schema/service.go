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
	Name   VarName        `json:"name"`
	Type   MethodArgument `json:"type"`   // TODO: VarType? Argument? ..perhaps ServiceArgument if we only allow struct..
	Stream bool           `json:"stream"` // TOOD(future)
}

type MethodOutput struct {
	Name   VarName        `json:"name"`
	Type   MethodArgument `json:"type"`   // TODO: same as input above..
	Stream bool           `json:"stream"` // TOOD(future)
}

type MethodArgument string
