package util

import "github.com/spf13/cobra"

func ChainRunE(fns ...func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		for _, fn := range fns {
			if fn == nil {
				continue
			}
			if err := fn(cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}
