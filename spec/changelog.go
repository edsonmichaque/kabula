package spec

type Changelog struct {
	Ref   string           `json:"ref,omitempty" xml:"ref,attr,omitempty"`
	Value []ChangelogValue `json:"value,omitempty" xml:"value,omitempty"`
}

type ChangelogValue struct {
	Title       string `json:"title,omitempty" xml:"title,omitempty"`
	Description string `json:"desc,omitempty" xml:"desc,omitempty"`
}
