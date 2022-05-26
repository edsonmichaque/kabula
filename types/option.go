package types

type Option struct {
	Ref   string       `json:"Ref,omitempty"`
	Value *OptionValue `json:"Value,omitempty"`
}

type OptionValue struct {
	Option string `json:"Option,omitempty"`
}
