package kpm

import (
	"github.com/edsonmichaque/kabula/archive"
	"github.com/spf13/cobra"
)

func CmdBuild() *cobra.Command {
	var flags struct {
		xml  bool
		sign bool
		zip  bool
		tar  bool
		gzip bool
	}

	cmd := &cobra.Command{
		Use:   "build",
		Short: "builds a kabula",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opt := archive.DefaultOptions
			if flags.xml {
				opt.Content = archive.FormatXML
			}

			if flags.zip {
				opt.Container = archive.Zip
			}

			if flags.tar {
				opt.Container = archive.Tar
			}

			if flags.gzip {
				opt.Container = archive.Gzip
			}

			return archive.Build(args[0], opt)
		},
	}

	cmd.Flags().BoolVarP(&flags.xml, "xml", "x", false, "")
	cmd.Flags().BoolVarP(&flags.sign, "sign", "s", false, "")
	cmd.Flags().BoolVarP(&flags.zip, "zip", "z", false, "zip")
	cmd.Flags().BoolVarP(&flags.gzip, "gzip", "g", false, "gzip")
	cmd.Flags().BoolVarP(&flags.tar, "tar", "t", false, "tar")

	return cmd
}
