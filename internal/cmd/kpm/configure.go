package kpm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdConfigure() *cobra.Command {
	var flags struct {
		debug   bool
		profile string
	}

	cmd := &cobra.Command{
		Use: "configure",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("configure")
			fmt.Println("debug:", flags.debug)

			return nil
		},
	}

	cmd.Flags().BoolVar(&flags.debug, "debug", false, "")
	cmd.Flags().StringVar(&flags.profile, "profile", "", "")

	return cmd
}
