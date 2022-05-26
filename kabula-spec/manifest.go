package spec

import "encoding/xml"

const (
	PkgExt = "kpm"
)

type Manifest struct {
	XMLName    xml.Name      `xml:"manifest"`
	Version    string        `json:"version,omitempty" xml:"version,omitempty"`
	Name       string        `json:"name,omitempty" xml:"name,omitempty"`
	Homepage   string        `json:"homepage,omitempty" xml:"homepage,omitempty"`
	License    string        `json:"license,omitempty" xml:"license,omitempty"`
	Publisher  Author        `json:"publisher,omitempty" xml:"publisher,omitempty"`
	Creator    Author        `json:"creator,omitempty" xml:"creator,omitempty"`
	Depends    []Requirement `json:"depends,omitempty" xml:"depends,omitempty"`
	Replaces   []Requirement `json:"replaces,omitempty" xml:"replaces,omitempty"`
	Recommends []Requirement `json:"recommends,omitempty" xml:"recommends,omitempty"`
	Suggests   []Requirement `json:"suggests,omitempty" xml:"suggests,omitempty"`
	Conflicts  []Requirement `json:"conflicts,omitempty" xml:"conflicts,omitempty"`
	Changelog  *Changelog    `json:"changelog,omitempty" xml:"changelog,omitempty"`
	Signatures *Signature    `json:"signatues,omitempty" xml:"sig,omitempty"`
}
