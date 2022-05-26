package kpm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdInstall() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "installs kabula package",
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("install")
			return nil
		},
	}

	return cmd
}
