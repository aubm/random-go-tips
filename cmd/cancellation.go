package cmd

import "github.com/spf13/cobra"

var cancellationCmd = &cobra.Command{
	Use:   "cancellation",
	Short: "A set of apps for demonstrating http request cancellation pitfalls",
}

func init() {
	cancellationCmd.AddCommand(cancellationUnhandledCmd)
	cancellationCmd.AddCommand(cancellationHandledCmd)
}
