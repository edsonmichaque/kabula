package kpm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdInfo() *cobra.Command {
	var flags struct {
		debug   bool
		profile string
		yes     bool
	}

	cmd := &cobra.Command{
		Use:  "info",
		Args: cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("info")
			return nil
		},
	}

	cmd.Flags().BoolVar(&flags.debug, "debug", false, "")
	cmd.Flags().BoolVarP(&flags.yes, "yes", "y", false, "")
	cmd.Flags().StringVar(&flags.profile, "profile", "", "")

	return cmd
}
