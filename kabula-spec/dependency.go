package spec

type Dependency struct {
	Ge string   `json:"ge,omitempty" xml:"ge,omitempty"`
	Le string   `json:"le,omitempty" xml:"le,omitempty"`
	Eq []string `json:"eq,omitempty" xml:"eq,omitempty"`
	Ne []string `json:"ne,omitempty" xml:"ne,omitempty"`
	Gt string   `json:"gt,omitempty" xml:"gt,omitempty"`
	Lt string   `json:"lt,omitempty" xml:"lt,omitempty"`
}
