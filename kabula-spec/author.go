package spec

type Author struct {
	Name  string `json:"name,omitempty" xml:"name,omitempty"`
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	URL   string `json:"url,omitempty" xml:"url,omitempty"`
}
