package kpm

import (
	"github.com/edsonmichaque/kabula/x/kab"
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
			opts := []kab.KabOption{}

			if flags.xml {
				opts = append(opts, kab.WithXML())
			}

			if flags.yaml {
				opts = append(opts, kab.WithYAML())
			}

			k := kab.NewKab(args[0], opts...)

			return k.Init()
		},
	}

	cmd.Flags().BoolVar(&flags.xml, "xml", false, "use XML format")
	cmd.Flags().BoolVar(&flags.json, "json", true, "use JSON format")
	cmd.Flags().BoolVar(&flags.yaml, "yaml", true, "use YAML format")

	return cmd
}
