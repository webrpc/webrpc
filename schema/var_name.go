package schema

type VarName string

// TODO:

// 1. Uppercase first letter
// 2. Downcase first letter

func (v *VarName) String() string {
	return string(*v)
}
