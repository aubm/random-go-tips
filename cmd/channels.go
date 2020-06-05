package cmd

import "github.com/spf13/cobra"

var channelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "A set of apps for demonstrating channels pitfalls",
}

func init() {
	channelsCmd.AddCommand(channelsSequenceCmd)
}
