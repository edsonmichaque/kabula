package kpm

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	spec "github.com/edsonmichaque/kabula/kabula-spec"
	"github.com/spf13/cobra"
)

func CmdInit() *cobra.Command {
	var flags struct {
		xml  bool
		json bool
	}

	cmd := &cobra.Command{
		Use:  "init",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := os.Mkdir(args[0], 0755); err != nil {
				if errors.Is(err, os.ErrExist) {
					return fmt.Errorf("%s already exist", args[0])
				}

				return err
			}

			ext := "json"
			if flags.xml {
				ext = "xml"
			}

			manifestPath := filepath.Join(args[0], fmt.Sprintf("manifest.%s", ext))
			f, err := os.Create(manifestPath)
			if err != nil {
				return err
			}

			defer func() {
				_ = f.Close()
			}()

			m := spec.Manifest{
				Version: "0.1.0",
				Name:    args[0],
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
			if flags.xml {
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
		},
	}

	cmd.Flags().BoolVar(&flags.xml, "xml", false, "use XML format")
	cmd.Flags().BoolVar(&flags.json, "json", true, "use JSON format")

	return cmd
}
