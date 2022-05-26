package kpm

import (
	"github.com/edsonmichaque/kabula/x/cli"
	"github.com/spf13/cobra"
)

var flags struct {
	config  string
	profile string
	debug   bool
}

func New(name string) error {
	cmd := cli.NewCommand(
		cli.WithHandler(func() *cobra.Command {

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
		cli.WithSubcommands(
			cli.NewCommand(cli.WithHandler(CmdConfigure)),
			cli.NewCommand(cli.WithHandler(CmdInstall)),
			cli.NewCommand(cli.WithHandler(CmdUpdate)),
			cli.NewCommand(cli.WithHandler(CmdRemove)),
			cli.NewCommand(cli.WithHandler(CmdList)),
			cli.NewCommand(cli.WithHandler(CmdSearch)),
			cli.NewCommand(cli.WithHandler(CmdFetch)),
			cli.NewCommand(cli.WithHandler(CmdBuild)),
			cli.NewCommand(cli.WithHandler(CmdInit)),
			cli.NewCommand(cli.WithHandler(CmdInfo)),
			cli.NewCommand(cli.WithHandler(CmdPush)),
			cli.NewCommand(
				cli.WithHandler(CmdRepo),
				cli.WithSubcommands(
					cli.NewCommand(cli.WithHandler(CmdRepoAdd)),
					cli.NewCommand(cli.WithHandler(CmdRepoRm)),
					cli.NewCommand(cli.WithHandler(CmdRepoLs)),
				),
			),
			cli.NewCommand(
				cli.WithHandler(CmdKey),
				cli.WithSubcommands(
					cli.NewCommand(cli.WithHandler(CmdKeyAdd)),
					cli.NewCommand(cli.WithHandler(CmdKeyRm)),
					cli.NewCommand(cli.WithHandler(CmdKeyLs)),
				),
			),
		),
	)

	return cmd.Execute()
}
