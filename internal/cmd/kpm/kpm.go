package kpm

import (
	cmdutil "github.com/edsonmichaque/kabula/kabula-cmdutil"
	"github.com/spf13/cobra"
)

func New(name string) error {
	cmd := cmdutil.New(
		cmdutil.WithHandler(func() *cobra.Command {
			return &cobra.Command{
				Use:           name,
				SilenceErrors: true,
				SilenceUsage:  true,
			}
		}),
		cmdutil.WithSubcommands(
			cmdutil.New(cmdutil.WithHandler(CmdConfigure)),
			cmdutil.New(cmdutil.WithHandler(CmdInstall)),
			cmdutil.New(cmdutil.WithHandler(CmdUpdate)),
			cmdutil.New(cmdutil.WithHandler(CmdRemove)),
			cmdutil.New(cmdutil.WithHandler(CmdList)),
			cmdutil.New(cmdutil.WithHandler(CmdSearch)),
			cmdutil.New(cmdutil.WithHandler(CmdFetch)),
			cmdutil.New(cmdutil.WithHandler(CmdBuild)),
			cmdutil.New(cmdutil.WithHandler(CmdInit)),
			cmdutil.New(cmdutil.WithHandler(CmdInfo)),
			cmdutil.New(cmdutil.WithHandler(CmdPush)),
		),
	)

	return cmd.Execute()
}
