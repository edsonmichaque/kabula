package kpm

import (
	"github.com/spf13/cobra"
)

func CmdKeyLs() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ls",
	}

	return cmd
}
