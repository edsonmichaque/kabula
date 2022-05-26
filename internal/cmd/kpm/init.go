package kpm

import (
	archive "github.com/edsonmichaque/kabula/archive"
	"github.com/spf13/cobra"
)

func CmdInit() *cobra.Command {
	var flags struct {
		xml     bool
		json    bool
		version string
	}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "create kabula package sources",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := archive.DefaultOptions

			if flags.xml {
				opts.Content = archive.FormatXML
			}

			return archive.New(args[0], opts)
		},
	}

	cmd.Flags().BoolVar(&flags.xml, "xml", false, "use XML format")
	cmd.Flags().BoolVar(&flags.json, "json", true, "use JSON format")
	cmd.Flags().StringVarP(&flags.version, "version", "V", "0.1.0", "version")

	return cmd
}
