package spec

import "encoding/xml"

const (
	Ext = "kar"
)

type Manifest struct {
	XMLName    xml.Name      `json:"-" yaml:"-" xml:"manifest"`
	Version    string        `json:"version,omitempty" yaml:"version" xml:"version,omitempty"`
	Name       string        `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Homepage   string        `json:"homepage,omitempty" yaml:"homepage,omitempty" xml:"homepage,omitempty"`
	License    string        `json:"license,omitempty" yaml:"license,omitempty" xml:"license,omitempty"`
	Publisher  *Author       `json:"publisher,omitempty" yaml:"publisher,omitempty" xml:"publisher,omitempty"`
	Creator    *Author       `json:"creator,omitempty" yaml:"creator,omitempty" xml:"creator,omitempty"`
	Depends    []Requirement `json:"depends,omitempty" yaml:"depends,omitempty" xml:"depends,omitempty"`
	Replaces   []Requirement `json:"replaces,omitempty" yaml:"replaces,omitempty" xml:"replaces,omitempty"`
	Recommends []Requirement `json:"recommends,omitempty" yaml:"recommends,omitempty" xml:"recommends,omitempty"`
	Suggests   []Requirement `json:"suggests,omitempty" yaml:"suggests,omitempty" xml:"suggests,omitempty"`
	Conflicts  []Requirement `json:"conflicts,omitempty" yaml:"conflicts,omitempty" xml:"conflicts,omitempty"`
	Changelog  *Changelog    `json:"changelog,omitempty" yaml:"changelog,omitempty" xml:"changelog,omitempty"`
	Signatures *Signature    `json:"signatues,omitempty" yaml:"signatures,omitempty" xml:"sig,omitempty"`
}
