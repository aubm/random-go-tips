package cmd

import (
	"github.com/aubm/random-go-tips/pkg/cmd/cancellation/unhandled"
	"github.com/spf13/cobra"
)

var cancellationUnhandledCmd = &cobra.Command{
	Use:   "unhandled --addr :8080",
	Short: "A web app that downloads and resizes images in go routines that are not cancelable",
	Run: func(cmd *cobra.Command, args []string) {
		unhandled.Run(globalConfig)
	},
}
