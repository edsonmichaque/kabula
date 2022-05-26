package kpm

import (
	"github.com/edsonmichaque/kabula/x/kab"
	"github.com/spf13/cobra"
)

func CmdBuild() *cobra.Command {
	var flags struct {
		xml  bool
		sign bool
		zip  bool
		tar  bool
		gzip bool
		zstd bool
	}

	cmd := &cobra.Command{
		Use:   "build",
		Short: "builds a kabula",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := []kab.KabOption{}

			if flags.xml {
				opts = append(opts, kab.WithXML())
			}

			if flags.zip {
				opts = append(opts, kab.WithZip())
			}

			if flags.gzip {
				opts = append(opts, kab.WithGZip())
			}

			if flags.zstd {
				opts = append(opts, kab.WithZStd())
			}

			k := kab.NewKab(args[0], opts...)

			return k.Init()
		},
	}

	cmd.Flags().BoolVarP(&flags.xml, "xml", "x", false, "")
	cmd.Flags().BoolVarP(&flags.sign, "sign", "s", false, "")
	cmd.Flags().BoolVarP(&flags.zip, "zip", "z", false, "zip")
	cmd.Flags().BoolVarP(&flags.gzip, "gzip", "g", false, "gzip")
	cmd.Flags().BoolVar(&flags.zstd, "zstd", false, "zstd")

	return cmd
}
