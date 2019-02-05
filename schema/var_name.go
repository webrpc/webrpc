package schema

import "strings"

type VarName string

// TitleDowncase will downcase the first letter of a string,
// ie. convert a name form 'FirstName' to 'firstName'
func (v VarName) TitleDowncase() string {
	if v == "" {
		return ""
	}
	s := string(v)
	return strings.ToLower(s[0:1]) + s[1:]
}

// TitleUpcase will upcase the first letter of a string,
// ie. convert a name form 'firstName' to 'FirstName'
func (v VarName) TitleUpcase() string {
	if v == "" {
		return ""
	}
	s := string(v)
	return strings.ToUpper(s[0:1]) + s[1:]
}

func (v VarName) Parse(schema *WebRPCSchema) error {
	return nil
}
