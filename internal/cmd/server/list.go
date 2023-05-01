package server

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/floshodan/hrobot-cli/internal/command"
	"github.com/floshodan/hrobot-cli/internal/state"
	"github.com/spf13/cobra"
)

func newListCommand(cli *state.State) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list FLAGS",
		Short:                 "List all Servers",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		//PreRunE:               util.ChainRunE(validateCreate, cli.EnsureToken),
		RunE: cli.Wrap(runList),
	}

	return cmd
}

func runList(cli *state.State, cmd *cobra.Command, args []string) error {
	servers, _, err := cli.Client().Server.List(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	table := &command.Table{
		Headers: []string{"ID", "ServerName", "ServerIP", "DataCenter", "Cancelled", "State"},
	}
	for _, server := range servers {
		serverstatus, _, _ := cli.Client().Reset.GetResetByServernumber(context.Background(), strconv.Itoa(server.ServerNumber))
		table.Rows = append(table.Rows, []string{
			strconv.Itoa(server.ServerNumber),
			server.ServerIP,
			server.ServerName,
			server.Dc,
			strconv.FormatBool(server.Cancelled),
			strings.ToUpper(serverstatus.OperatingStatus),
		})
	}
	table.Renderer()

	return nil
}
