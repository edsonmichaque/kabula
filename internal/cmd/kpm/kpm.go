package kpm

import (
	"github.com/edsonmichaque/kabula/cmdutil"
	"github.com/spf13/cobra"
)

func New(name string) error {
	cmd := cmdutil.NewCommand(
		cmdutil.WithHandler(func() *cobra.Command {
			return &cobra.Command{
				Use: name,
			}
		}),
		cmdutil.WithChildren(
			cmdutil.NewCommand(cmdutil.WithHandler(CmdConfigure)),
			cmdutil.NewCommand(cmdutil.WithHandler(CmdInstall)),
			cmdutil.NewCommand(cmdutil.WithHandler(CmdUpdate)),
			cmdutil.NewCommand(cmdutil.WithHandler(CmdRemove)),
			cmdutil.NewCommand(cmdutil.WithHandler(CmdList)),
			cmdutil.NewCommand(cmdutil.WithHandler(CmdSearch)),
			cmdutil.NewCommand(cmdutil.WithHandler(CmdFetch)),
			cmdutil.NewCommand(cmdutil.WithHandler(CmdBuild)),
			cmdutil.NewCommand(cmdutil.WithHandler(CmdInit)),
		),
	)

	return cmd.Execute()
}
