package kpm

import (
	"github.com/spf13/cobra"
)

func CmdKey() *cobra.Command {
	cmd := &cobra.Command{
		Use: "key",
	}

	return cmd
}
