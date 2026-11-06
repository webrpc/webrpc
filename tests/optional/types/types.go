package types

type ZeroValue string

func (v ZeroValue) IsZero() bool { return v == "" }

type PointerZero string

func (v *PointerZero) IsZero() bool { return *v == "" }
