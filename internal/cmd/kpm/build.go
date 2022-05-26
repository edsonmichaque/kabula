package kpm

import (
	"archive/tar"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	spec "github.com/edsonmichaque/kabula/kabula-spec"
	"github.com/spf13/cobra"
)

func CmdBuild() *cobra.Command {
	var flags struct {
		debug   bool
		profile string
		yes     bool
	}

	cmd := &cobra.Command{
		Use:  "build",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fullpath := filepath.Join(args[0], "manifest.json")
			f, err := os.Open(fullpath)
			if err != nil {
				return err
			}

			defer func() {
				_ = f.Close()
			}()

			data, err := io.ReadAll(f)
			if err != nil {
				return err
			}

			var m spec.Manifest
			if err := json.Unmarshal(data, &m); err != nil {
				return err
			}

			tarball, err := os.Create(fmt.Sprintf("%s_%s.%s", m.Name, m.Version, spec.PkgExt))
			if err != nil {
				return err
			}

			dir, err := os.Open(args[0])
			if err != nil {
				return err
			}

			l, err := dir.Readdir(0)
			if err != nil {
				return err
			}

			t := tar.NewWriter(tarball)
			defer t.Close()
			for _, e := range l {
				if e.IsDir() {
					continue
				}

				c, err := os.Open(filepath.Join(args[0], e.Name()))
				if err != nil {
					return err
				}

				if err := t.WriteHeader(&tar.Header{
					Name:     strings.TrimPrefix(c.Name(), fmt.Sprintf("%s%c", args[0], filepath.Separator)),
					Size:     e.Size(),
					Mode:     int64(e.Mode()),
					ModTime:  e.ModTime(),
					Typeflag: tar.TypeReg,
				}); err != nil {
					return err
				}

				if _, err := io.Copy(t, c); err != nil {
					return err
				}

				c.Close()
			}

			fmt.Println(tarball.Name())

			return nil
		},
	}

	cmd.Flags().BoolVar(&flags.debug, "debug", false, "")
	cmd.Flags().BoolVarP(&flags.yes, "yes", "y", false, "")
	cmd.Flags().StringVar(&flags.profile, "profile", "", "")

	return cmd
}
