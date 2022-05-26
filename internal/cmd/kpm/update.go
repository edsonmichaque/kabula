package kpm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdUpdate() *cobra.Command {
	var flags struct {
		debug   bool
		profile string
		yes     bool
	}

	cmd := &cobra.Command{
		Use:  "update",
		Args: cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("update")
			return nil
		},
	}

	cmd.Flags().BoolVar(&flags.debug, "debug", false, "")
	cmd.Flags().BoolVarP(&flags.yes, "yes", "y", false, "")
	cmd.Flags().StringVar(&flags.profile, "profile", "", "")

	return cmd
}
