package cmd

import (
	"github.com/aubm/random-go-tips/pkg/cmd/cancellation/handled"
	"github.com/spf13/cobra"
)

var cancellationHandledCmd = &cobra.Command{
	Use:   "handled --addr :8080",
	Short: "A web app that downloads and resizes images in cancelable go routines",
	Run: func(cmd *cobra.Command, args []string) {
		handled.Run(globalConfig)
	},
}
