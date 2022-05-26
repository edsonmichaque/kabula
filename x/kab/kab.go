package kab

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type archive int

const (
	formatZip archive = iota
	formatGZip
	formatZStd
)

func (c archive) String() string {
	return [...]string{"zip", "gzip", "zstd"}[c]
}

type content int

const (
	formatJSON content = iota
	formatXML
	formatYAML
)

func (c content) Ext() string {
	return [...]string{"json", "xml", "yml"}[c]
}

type KabOption func(*Kab)

type Kab struct {
	archive archive
	content content
	dst     string
	src     string
}

func NewKab(target string, opts ...KabOption) *Kab {
	kab := &Kab{
		src: target,
	}

	for _, opt := range opts {
		opt(kab)
	}

	return kab
}

func WithJSON() KabOption {
	return func(k *Kab) {
		k.content = formatJSON
	}
}

func WithXML() KabOption {
	return func(k *Kab) {
		k.content = formatXML
	}
}

func WithYAML() KabOption {
	return func(k *Kab) {
		k.content = formatYAML
	}
}

func WithZip() KabOption {
	return func(k *Kab) {
		k.archive = formatZip
	}
}

func WithGZip() KabOption {
	return func(k *Kab) {
		k.archive = formatGZip
	}
}

func WithZStd() KabOption {
	return func(k *Kab) {
		k.archive = formatZStd
	}
}

func (k Kab) Marshal(src interface{}) ([]byte, error) {
	if k.content == formatXML {
		return xml.MarshalIndent(src, "", "  ")
	}

	if k.content == formatYAML {
		return yaml.Marshal(src)
	}

	return json.MarshalIndent(src, "", "  ")
}

func (k Kab) ManifestPath() string {
	return filepath.Join(k.src, fmt.Sprintf("manifest.%s", k.content.Ext()))
}

func (k Kab) Build() error {
	if k.archive == formatZip {

	}

	if k.archive == formatGZip {

	}

	if k.archive == formatZStd {

	}

	return errors.New("unknown format")
}
