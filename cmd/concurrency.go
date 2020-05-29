package cmd

import "github.com/spf13/cobra"

var concurrencyCmd = &cobra.Command{
	Use:   "concurrency",
	Short: "A set of apps for demonstrating concurrency caveats",
}

func init() {
	concurrencyCmd.AddCommand(concurrencySequenceCmd)
	concurrencyCmd.AddCommand(concurrencyUnboundCmd)
}
