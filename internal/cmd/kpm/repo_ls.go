package kpm

import (
	"github.com/spf13/cobra"
)

func CmdRepoLs() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ls",
	}

	return cmd
}
