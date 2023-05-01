package sshkey

import (
	"fmt"

	"github.com/floshodan/hrobot-cli/internal/state"
	"github.com/spf13/cobra"
)

func newPrintCommand(cli *state.State) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "print FLAGS",
		Short:                 "Print PublicKey to stdo",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		//PreRunE:               util.ChainRunE(validateCreate, cli.EnsureToken),
		RunE: cli.Wrap(runPrint),
	}

	return cmd
}

func runPrint(cli *state.State, cmd *cobra.Command, args []string) error {
	sshKeys, _, err := cli.Client().SSHKey.List(cli.Context)
	if err != nil {
		return nil
	}
	if len(args) < 1 {
		fmt.Println("Name needs to be defined")
		return nil
	}
	name := args[0]
	var data string
	for _, sshKey := range sshKeys {
		if sshKey.Name == name {
			data = sshKey.Data
			break
		}
	}
	if data == "" {
		fmt.Printf("Key with name '%s' not found \n", name)
	} else {
		fmt.Println(data)
	}

	return nil
}
