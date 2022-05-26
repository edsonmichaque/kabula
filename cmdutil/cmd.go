package cmdutil

import "github.com/spf13/cobra"

type Command struct {
	Handler  *cobra.Command
	Children []*Command
}

func (c *Command) apply() {
	for _, child := range c.Children {
		if child.Children != nil {
			child.apply()
		}

		c.Handler.AddCommand(child.Handler)
	}
}

func (c *Command) Execute() error {
	c.apply()
	return c.Handler.Execute()
}

type CommandOption func(*Command)

func WithHandler(f func() *cobra.Command) CommandOption {
	return func(c *Command) {
		c.Handler = f()
	}
}

func WithChildren(sub ...*Command) CommandOption {
	return func(c *Command) {
		if c.Children == nil {
			c.Children = make([]*Command, 0)
		}

		c.Children = append(c.Children, sub...)
	}

}

func NewCommand(opts ...CommandOption) *Command {
	c := &Command{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
