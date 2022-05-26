package archive

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/edsonmichaque/kabula/spec"
)

func New(target string, args Options) error {
	if err := os.Mkdir(target, 0755); err != nil {
		if errors.Is(err, os.ErrExist) {
			return fmt.Errorf("%s already exist", target)
		}

		return err
	}

	ext := "json"
	if args.Content == FormatXML {
		ext = "xml"
	}

	manifestPath := filepath.Join(target, fmt.Sprintf("manifest.%s", ext))
	f, err := os.Create(manifestPath)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	m := spec.Manifest{
		Version: "0.1.0",
		Name:    target,
		Depends: []spec.Requirement{
			{
				Package: "example.com/foo",
				Version: spec.Dependency{
					Ge: "1.0.0",
				},
			},
			{
				Package: "example.com/bar",
				Version: spec.Dependency{
					Eq: []string{"1.0.0", "1.0.1"},
				},
			},
		},
		Publisher: spec.Author{
			Name:  "Edson Michaque",
			Email: "edson@michaque.com",
		},
		Creator: spec.Author{
			Name:  "Edson Michaque",
			Email: "edson@michaque.com",
		},
		Signatures: &spec.Signature{
			Ref: "sig.json",
		},
		Changelog: &spec.Changelog{
			Ref: "changelog.json",
		},
	}

	var data []byte
	if args.Content == FormatXML {
		data, err = xml.MarshalIndent(m, "", "    ")
		if err != nil {
			return err
		}
	} else {
		data, err = json.MarshalIndent(m, "", "    ")
		if err != nil {
			return err
		}
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}
