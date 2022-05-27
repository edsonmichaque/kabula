package kpm

import (
	"github.com/edsonmichaque/kabula/x/kar"
	"github.com/spf13/cobra"
)

func CmdBuild() *cobra.Command {
	var flags struct {
		xml     bool
		sign    bool
		zip     bool
		tar     bool
		gzip    bool
		zstd    bool
		xz      bool
		verbose bool
	}

	cmd := &cobra.Command{
		Use:   "build",
		Short: "builds a kabula",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := []kar.KabOption{}

			if flags.zip {
				opts = append(opts, kar.WithZip())
			}

			if flags.gzip {
				opts = append(opts, kar.WithGZip())
			}

			if flags.zstd {
				opts = append(opts, kar.WithZStd())
			}

			if flags.xz {
				opts = append(opts, kar.WithXZ())
			}

			if flags.verbose {
				opts = append(opts, kar.WithLogs())
			}

			k := kar.New(args[0], opts...)

			return k.Build()
		},
	}

	cmd.Flags().BoolVarP(&flags.xml, "xml", "x", false, "")
	cmd.Flags().BoolVarP(&flags.sign, "sign", "s", false, "")
	cmd.Flags().BoolVarP(&flags.zip, "zip", "z", false, "zip")
	cmd.Flags().BoolVarP(&flags.gzip, "gzip", "g", false, "gzip")
	cmd.Flags().BoolVar(&flags.zstd, "zstd", false, "zstd")
	cmd.Flags().BoolVar(&flags.xz, "xz", false, "xz")
	cmd.Flags().BoolVar(&flags.verbose, "verbose", false, "verbose")

	return cmd
}
