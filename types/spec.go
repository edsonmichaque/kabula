package types

import "encoding/xml"

const (
	PkgExt = "kpm"
)

// Manifest.json
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
	Sig        *Sig          `json:"sig,omitempty" xml:"sig,omitempty"`
}

type Requirement struct {
	Package string     `json:"package,omitempty" xml:"package,omitempty"`
	Version Dependency `json:"version,omitempty" xml:"version,omitempty"`
}

type Dependency struct {
	Ge string   `json:"ge,omitempty" xml:"ge,omitempty"`
	Le string   `json:"le,omitempty" xml:"le,omitempty"`
	Eq []string `json:"eq,omitempty" xml:"eq,omitempty"`
	Ne []string `json:"ne,omitempty" xml:"ne,omitempty"`
	Gt string   `json:"gt,omitempty" xml:"gt,omitempty"`
	Lt string   `json:"lt,omitempty" xml:"lt,omitempty"`
}

// Sig.json
type Sig struct {
	Ref   string    `json:"ref,omitempty" xml:"ref,attr,omitempty"`
	Value *SigValue `json:"value,omitempty" xml:",omitempty"`
}

type SigValue struct {
	MD5    string  `json:"md5,omitempty" xml:"md5,omitempty"`
	Sha1   string  `json:"sha1,omitempty" xml:"sha1,omitempty"`
	Sha256 string  `json:"sha256,omitempty" xml:"sha256,omitempty"`
	Sha512 string  `json:"sha512,omitempty" xml:"sha512,omitempty"`
	GPG    *GPGSig `json:"gpg,omitempty" xml:"gpg,omitempty"`
}

type GPGSig struct {
	Fingerprint string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
	Signature   string `json:"signature,omitempty" xml:"signature,omitempty"`
}

// Changelog.json
type Changelog struct {
	Ref   string           `json:"ref,omitempty" xml:"ref,attr,omitempty"`
	Value []ChangelogValue `json:"value,omitempty" xml:"value,omitempty"`
}

type ChangelogValue struct {
	Title       string `json:"title,omitempty" xml:"title,omitempty"`
	Description string `json:"desc,omitempty" xml:"desc,omitempty"`
}

type Author struct {
	Name  string `json:"name,omitempty" xml:"name,omitempty"`
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	URL   string `json:"url,omitempty" xml:"url,omitempty"`
}
