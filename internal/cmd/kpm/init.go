package kpm

import (
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
		answersDir   bool
		questionsDir bool
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

			manifestPath := filepath.Join(args[0], "manifest.json")
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

			data, err := xml.MarshalIndent(m, "", "    ")
			if err != nil {
				return err
			}

			if _, err := f.Write(data); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&flags.answersDir, "answers-dir", false, "answers dir")
	cmd.Flags().BoolVar(&flags.questionsDir, "questions-dir", false, "answers dir")

	return cmd
}
