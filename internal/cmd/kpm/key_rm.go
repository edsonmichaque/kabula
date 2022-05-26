package kpm

import (
	"github.com/spf13/cobra"
)

func CmdKeyRm() *cobra.Command {
	var flags struct {
		all   bool
		force bool
	}

	cmd := &cobra.Command{
		Use: "rm",
	}

	cmd.Flags().BoolVarP(&flags.all, "all", "a", false, "remove all")
	cmd.Flags().BoolVarP(&flags.force, "force", "f", false, "force")

	return cmd
}
