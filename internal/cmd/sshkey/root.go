package sshkey

import (
	"github.com/floshodan/hrobot-cli/internal/hrapi"
	"github.com/floshodan/hrobot-cli/internal/state"
	"github.com/spf13/cobra"
)

func NewCommand(cli *state.State, client hrapi.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "ssh-key",
		Short:                 "Manage SSH keys",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		//listCmd.CobraCommand(cli.Context, client, cli),
		newListCommand(cli),
		newCreateCommand(cli),
		newPrintCommand(cli),
		//updateCmd.CobraCommand(cli.Context, client, cli),
		//deleteCmd.CobraCommand(cli.Context, client, cli),
		//describeCmd.CobraCommand(cli.Context, client, cli),
		//labelCmds.AddCobraCommand(cli.Context, client, cli),
		//labelCmds.RemoveCobraCommand(cli.Context, client, cli),
	)
	return cmd
}
