package kpm

import (
	"fmt"

	"github.com/edsonmichaque/kabula/x/kar"
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
			opts, err := kar.Info(args[0])
			if err != nil {
				return err
			}

			fmt.Printf("%#v\n", opts.Container.String())

			return nil
		},
	}

	cmd.Flags().BoolVar(&flags.xml, "xml", false, "xml")

	return cmd
}
