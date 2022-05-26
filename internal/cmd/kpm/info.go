package kpm

import (
	"fmt"

	"github.com/edsonmichaque/kabula/x/kab"
	"github.com/spf13/cobra"
)

func CmdInfo() *cobra.Command {
	var flags struct {
		xml bool
	}

	cmd := &cobra.Command{
		Use:   "info",
		Short: "installs kabula package",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts, err := kab.Info(args[0])
			if err != nil {
				return err
			}

			if opts.Container == kab.Tar {
				fmt.Println("container format tar")
			}

			if opts.Container == kab.Zip {
				fmt.Println("container format zip")
			}

			if opts.Container == kab.Gzip {
				fmt.Println("container format gzip")
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&flags.xml, "xml", false, "xml")

	return cmd
}
