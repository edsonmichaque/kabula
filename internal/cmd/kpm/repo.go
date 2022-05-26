package kpm

import (
	"github.com/spf13/cobra"
)

func CmdRepo() *cobra.Command {
	cmd := &cobra.Command{
		Use: "repo",
	}

	return cmd
}
