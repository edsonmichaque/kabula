package spec

type Requirement struct {
	Package string     `json:"package,omitempty" xml:"package,omitempty"`
	Version Dependency `json:"version,omitempty" xml:"version,omitempty"`
}
