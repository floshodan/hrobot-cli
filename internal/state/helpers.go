package state

import (
	"errors"

	"github.com/spf13/cobra"
)

func (c *State) Wrap(f func(*State, *cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return f(c, cmd, args)
	}
}

func (c *State) EnsureToken(cmd *cobra.Command, args []string) error {
	if c.Token == "" {
		return errors.New("no active context or token (see `hrobot context --help`)")
	}
	return nil
}
