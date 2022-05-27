package kar

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/edsonmichaque/kabula/x/kar/builder.go"
	"github.com/edsonmichaque/kabula/x/spec"
	"gopkg.in/yaml.v3"
)

type archiveFormat int

const (
	Zip archiveFormat = iota + 1
	GZip
	ZStd
	XZ
	InvalidArchive
)

func (c archiveFormat) String() string {
	return map[archiveFormat]string{
		Zip:            "zip",
		GZip:           "gzip",
		ZStd:           "zstd",
		XZ:             "xz",
		InvalidArchive: "invalid",
	}[c]
}

type content int

const (
	JSON content = iota
	XML
	YAML
	Unknown
)

func (c content) Ext() string {
	return [...]string{"json", "xml", "yml"}[c]
}

type KabOption func(*Kab)

type Kab struct {
	archive archiveFormat
	content content
	dst     string
	src     string
	log     bool
}

func New(target string, opts ...KabOption) *Kab {
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
		k.content = JSON
	}
}

func WithXML() KabOption {
	return func(k *Kab) {
		k.content = XML
	}
}

func WithYAML() KabOption {
	return func(k *Kab) {
		k.content = YAML
	}
}

func WithZip() KabOption {
	return func(k *Kab) {
		k.archive = Zip
	}
}

func WithGZip() KabOption {
	return func(k *Kab) {
		k.archive = GZip
	}
}

func WithZStd() KabOption {
	return func(k *Kab) {
		k.archive = ZStd
	}
}

func WithXZ() KabOption {
	return func(k *Kab) {
		k.archive = XZ
	}
}

func WithLogs() KabOption {
	return func(k *Kab) {
		k.log = true
	}
}

func (k Kab) decode() (*spec.Manifest, error) {
	fullpath := filepath.Join(k.src, fmt.Sprintf("%s.%s", "manifest", k.content.Ext()))
	f, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var manifest spec.Manifest

	if k.content == XML {
		if err := xml.Unmarshal(data, &manifest); err != nil {
			return nil, err
		}
	}

	if k.content == JSON {
		if err := json.Unmarshal(data, &manifest); err != nil {
			return nil, err
		}
	}

	if k.content == YAML {
		if err := yaml.Unmarshal(data, &manifest); err != nil {
			return nil, err
		}
	}

	return &manifest, nil
}

func (k Kab) encode(src interface{}) ([]byte, error) {
	if k.content == XML {
		return xml.MarshalIndent(src, "", "  ")
	}

	if k.content == YAML {
		return yaml.Marshal(src)
	}

	return json.MarshalIndent(src, "", "  ")
}

func (k Kab) ManifestPath() string {
	return filepath.Join(k.src, fmt.Sprintf("manifest.%s", k.content.Ext()))
}

func (k Kab) create(manifest *spec.Manifest) (*os.File, error) {
	return os.Create(fmt.Sprintf("%s_%s.%s", manifest.Name, manifest.Version, spec.Ext))
}

func (k *Kab) walk(entries []builder.Entry) error {
	c := Unknown
	for _, e := range entries {
		m := formats()

		if content, ok := m[e.Name]; ok {
			c = content
			break
		}
	}

	if k.content == Unknown {
		return fmt.Errorf("dir does not have a manifest file")
	}

	k.content = c

	return nil
}

func formats() map[string]content {
	return map[string]content{
		jsonManifest: JSON,
		xmlManifest:  XML,
		yamlManifest: YAML,
	}
}

func builders() map[archiveFormat]builder.Builder {
	return map[archiveFormat]builder.Builder{
		GZip: builder.TGz,
		Zip:  builder.Zip,
		ZStd: builder.ZStd,
		XZ:   builder.TXz,
	}
}

const (
	jsonManifest = "manifest.json"
	xmlManifest  = "manifest.xml"
	yamlManifest = "manifest.yml"
)
