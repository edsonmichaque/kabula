package kpm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdPush() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "push",
		Args: cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("push")
			return nil
		},
	}

	return cmd
}
