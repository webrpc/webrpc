package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequiredFields(t *testing.T) {
	tests := []struct {
		name   string
		fields []*TypeField
		want   []*TypeField
	}{
		{
			name: "one field is required",
			fields: []*TypeField{
				{
					Name: "Name",
					Type: &VarType{
						Type: T_String,
					},
					TypeExtra: TypeExtra{Optional: false},
				},
				{
					Name: "Phone number",
					Type: &VarType{
						Type: T_String,
					},
					TypeExtra: TypeExtra{Optional: true},
				},
			},
			want: []*TypeField{
				{
					Name: "Name",
					Type: &VarType{
						Type: T_String,
					},
					TypeExtra: TypeExtra{Optional: false},
				},
			},
		},
		{
			name: "all fields are optional",
			fields: []*TypeField{
				{
					Name: "Name",
					Type: &VarType{
						Type: T_String,
					},
					TypeExtra: TypeExtra{Optional: true},
				},
			},
			want: []*TypeField{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ty := &Type{Fields: tt.fields}
			assert.Equal(t, tt.want, ty.RequiredFields())
		})
	}
}
