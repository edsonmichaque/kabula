package cli

import (
	"github.com/spf13/cobra"
)

type Option func(*Command)

type Command struct {
	handler     *cobra.Command
	subcommands []*Command
}

func (c *Command) populateSubcommands() {
	for _, child := range c.subcommands {
		if child.subcommands != nil {
			child.populateSubcommands()
		}
		c.handler.AddCommand(child.handler)
	}
}

func (c *Command) Execute() error {
	c.populateSubcommands()
	return c.handler.Execute()
}

func WithHandler(f func() *cobra.Command) Option {
	return func(c *Command) {
		c.handler = f()
	}
}

func WithSubcommands(sub ...*Command) Option {
	return func(c *Command) {
		if c.subcommands == nil {
			c.subcommands = make([]*Command, 0)
		}
		c.subcommands = append(c.subcommands, sub...)
	}
}

func NewCommand(opts ...Option) *Command {
	c := &Command{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
