package kpm

import (
	"github.com/edsonmichaque/kabula/x/kar"
	"github.com/spf13/cobra"
)

func CmdInit() *cobra.Command {
	var flags struct {
		xml  bool
		json bool
		yaml bool
	}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "create kabula package sources",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := []kar.KabOption{}

			if flags.xml {
				opts = append(opts, kar.WithXML())
			}

			if flags.yaml {
				opts = append(opts, kar.WithYAML())
			}

			k := kar.New(args[0], opts...)

			return k.Init()
		},
	}

	cmd.Flags().BoolVarP(&flags.xml, "use-xml", "x", false, "use XML format")
	cmd.Flags().BoolVarP(&flags.json, "use-json", "j", false, "use JSON format")
	cmd.Flags().BoolVarP(&flags.yaml, "use-yaml", "Y", false, "use YAML format")

	return cmd
}
