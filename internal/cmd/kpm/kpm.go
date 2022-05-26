package kpm

import (
	"github.com/edsonmichaque/libcmd"
	"github.com/spf13/cobra"
)

func New(name string) error {
	cmd := libcmd.New(
		libcmd.WithHandler(func() *cobra.Command {
			return &cobra.Command{
				Use: name,
			}
		}),
		libcmd.WithSubcommands(
			libcmd.New(libcmd.WithHandler(CmdConfigure)),
			libcmd.New(libcmd.WithHandler(CmdInstall)),
			libcmd.New(libcmd.WithHandler(CmdUpdate)),
			libcmd.New(libcmd.WithHandler(CmdRemove)),
			libcmd.New(libcmd.WithHandler(CmdList)),
			libcmd.New(libcmd.WithHandler(CmdSearch)),
			libcmd.New(libcmd.WithHandler(CmdFetch)),
			libcmd.New(libcmd.WithHandler(CmdBuild)),
			libcmd.New(libcmd.WithHandler(CmdInit)),
			libcmd.New(libcmd.WithHandler(CmdInfo)),
		),
	)

	return cmd.Execute()
}
