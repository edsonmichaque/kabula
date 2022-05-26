package kpm

import (
	"github.com/edsonmichaque/kabula/archive"
	"github.com/spf13/cobra"
)

func CmdBuild() *cobra.Command {
	var flags struct {
		xml   bool
		sign  bool
		zip   bool
		tar   bool
		gzip  bool
		bzip2 bool
	}

	cmd := &cobra.Command{
		Use:   "build",
		Short: "builds a kabula",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			archiveArgs := archive.DefaultArgs

			if flags.xml {
				archiveArgs.Content = archive.FormatXML
			}

			if flags.zip {
				archiveArgs.Container = archive.KindZip
			}

			return archive.Build(args[0], archiveArgs)
		},
	}

	cmd.Flags().BoolVarP(&flags.xml, "xml", "x", false, "")
	cmd.Flags().BoolVarP(&flags.sign, "sign", "s", false, "")
	cmd.Flags().BoolVarP(&flags.zip, "zip", "z", false, "zip")
	cmd.Flags().BoolVarP(&flags.gzip, "gzip", "g", false, "gzip")
	cmd.Flags().BoolVarP(&flags.bzip2, "bzip2", "b", false, "bzip2")

	return cmd
}
