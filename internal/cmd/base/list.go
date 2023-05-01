package base

import (
	"context"

	"github.com/floshodan/hrobot-cli/internal/cmd/output"
	"github.com/floshodan/hrobot-cli/internal/hrapi"
	"github.com/spf13/cobra"
)

type ListCmd struct {
	ResourceNamePlural string // e.g. "servers"
	DefaultColumns     []string
	Fetch              func(context.Context, hrapi.Client, *cobra.Command, []string) ([]interface{}, error)
	AdditionalFlags    func(*cobra.Command)
	OutputTable        func(client hrapi.Client) *output.Table
	JSONSchema         func([]interface{}) interface{}
}

/*
// CobraCommand creates a command that can be registered with cobra.
func (lc *ListCmd) CobraCommand(
	ctx context.Context, client hrapi.Client,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list [FlAGS]",
		Short: fmt.Sprintf("List %s", lc.ResourceNamePlural),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               tokenEnsurer.EnsureToken,
		RunE: func(cmd *cobra.Command, args []string) error {
			return lc.Run(ctx, client, cmd)
		},
	}
	output.AddFlag(cmd, output.OptionNoHeader(), output.OptionColumns(outputColumns), output.OptionJSON())
	cmd.Flags().StringP("selector", "l", "", "Selector to filter by labels")
	if lc.AdditionalFlags != nil {
		lc.AdditionalFlags(cmd)
	}
	cmd.Flags().StringSliceP("sort", "s", []string{"id:asc"}, "Determine the sorting of the result")
	return cmd
}

// Run executes a list command
func (lc *ListCmd) Run(ctx context.Context, client hcapi2.Client, cmd *cobra.Command) error {
	outOpts := output.FlagsForCommand(cmd)

	labelSelector, _ := cmd.Flags().GetString("selector")
	listOpts := hcloud.ListOpts{
		LabelSelector: labelSelector,
		PerPage:       50,
	}
	sorts, _ := cmd.Flags().GetStringSlice("sort")

	resources, err := lc.Fetch(ctx, client, cmd, listOpts, sorts)
	if err != nil {
		return err
	}

	if outOpts.IsSet("json") {
		return util.DescribeJSON(lc.JSONSchema(resources))
	}

	cols := lc.DefaultColumns
	if outOpts.IsSet("columns") {
		cols = outOpts["columns"]
	}

	table := lc.OutputTable(client)
	if !outOpts.IsSet("noheader") {
		table.WriteHeader(cols)
	}
	for _, resource := range resources {
		table.Write(cols, resource)
	}
	table.Flush()

	return nil
}
*/
