package kpm

import (
	cmdutil "github.com/edsonmichaque/kabula/x/cmdutil"
	"github.com/spf13/cobra"
)

var flags struct {
	config  string
	profile string
	debug   bool
}

func New(name string) error {
	cmd := cmdutil.New(
		cmdutil.WithHandler(func() *cobra.Command {

			cmd := &cobra.Command{
				Use:           name,
				Version:       "0.1.0",
				SilenceErrors: true,
				SilenceUsage:  true,
			}

			cmd.PersistentFlags().StringVarP(&flags.config, "config", "c", "", "configuration file")
			cmd.PersistentFlags().StringVarP(&flags.profile, "profile", "p", "", "profile")
			cmd.PersistentFlags().BoolVarP(&flags.debug, "debug", "D", false, "debug")

			return cmd
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
			cmdutil.New(
				cmdutil.WithHandler(CmdRepo),
				cmdutil.WithSubcommands(
					cmdutil.New(cmdutil.WithHandler(CmdRepoAdd)),
					cmdutil.New(cmdutil.WithHandler(CmdRepoRm)),
					cmdutil.New(cmdutil.WithHandler(CmdRepoLs)),
				),
			),
			cmdutil.New(
				cmdutil.WithHandler(CmdKey),
				cmdutil.WithSubcommands(
					cmdutil.New(cmdutil.WithHandler(CmdKeyAdd)),
					cmdutil.New(cmdutil.WithHandler(CmdKeyRm)),
					cmdutil.New(cmdutil.WithHandler(CmdKeyLs)),
				),
			),
		),
	)

	return cmd.Execute()
}
