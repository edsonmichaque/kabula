package kpm

import (
	"github.com/spf13/cobra"
)

func CmdRepoAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add",
	}

	return cmd
}
