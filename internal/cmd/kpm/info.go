package kpm

import (
	"github.com/edsonmichaque/kabula/archive"
	"github.com/spf13/cobra"
)

func CmdInfo() *cobra.Command {
	var flags struct {
		debug   bool
		profile string
		yes     bool
	}

	cmd := &cobra.Command{
		Use:   "info",
		Short: "installs kabula package",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return archive.Info(args[0])
		},
	}

	cmd.Flags().BoolVar(&flags.debug, "debug", false, "")
	cmd.Flags().BoolVarP(&flags.yes, "yes", "y", false, "")
	cmd.Flags().StringVar(&flags.profile, "profile", "", "")

	return cmd
}
