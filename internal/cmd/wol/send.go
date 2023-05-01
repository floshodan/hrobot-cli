package wol

import (
	"fmt"
	"log"

	"github.com/floshodan/hrobot-cli/internal/state"
	"github.com/spf13/cobra"
)

func newSendCommand(cli *state.State) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "send [",
		Short:                 "Send a WOL Packet to server",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		//PreRunE:               util.ChainRunE(validateCreate, cli.EnsureToken),
		RunE: cli.Wrap(runSend),
	}
	cmd.Flags().String("name", "", "Key name (required)")
	return cmd
}

func runSend(cli *state.State, cmd *cobra.Command, args []string) error {
	//cli.Client().WakeOnLane
	if len(args) < 1 {
		fmt.Println("Servernumber or ServerIP must be provided")
		return nil
	}
	servernumber := args[0]
	_, resp, err := cli.Client().WakeOnLane.Send(cli.Context, servernumber)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == 404 {
		fmt.Println("Servernumber or IP not found")
	}
	if resp.StatusCode == 200 {
		fmt.Printf("Wake on Lan packet send to %s \n", servernumber)
	}
	return nil
}
