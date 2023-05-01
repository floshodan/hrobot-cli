package server

import (
	"github.com/floshodan/hrobot-cli/internal/hrapi"
	"github.com/floshodan/hrobot-cli/internal/state"
	"github.com/spf13/cobra"
)

func NewCommand(cli *state.State, client hrapi.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "server",
		Short:                 "Mange Servers",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		newListCommand(cli),
	)

	return cmd
}
