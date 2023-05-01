package sshkey

import (
	"strconv"

	"github.com/floshodan/hrobot-cli/internal/command"
	"github.com/floshodan/hrobot-cli/internal/state"
	"github.com/spf13/cobra"
)

type ListCommand struct {
}

func newListCommand(cli *state.State) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list FLAGS",
		Short:                 "List all Keys",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		//PreRunE:               util.ChainRunE(validateCreate, cli.EnsureToken),
		RunE: cli.Wrap(runList),
	}

	return cmd
}

func runList(cli *state.State, cmd *cobra.Command, args []string) error {
	sshKeys, _, err := cli.Client().SSHKey.List(cli.Context)
	if err != nil {
		return nil
	}
	table := &command.Table{
		Headers: []string{"Name", "Fingerprint", "Type", "Size"},
	}
	for _, sshKey := range sshKeys {
		table.Rows = append(table.Rows, []string{sshKey.Name, sshKey.Fingerprint, sshKey.Type, strconv.Itoa(sshKey.Size)})
	}
	table.Renderer()
	//renderSSHKeyTable(sshKeys)

	return nil
}
