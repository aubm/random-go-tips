package cmd

import (
	"github.com/spf13/cobra"

	"github.com/aubm/random-go-tips/pkg/cmd/channels/sequence"
)

var channelsSequenceCmd = &cobra.Command{
	Use:   "sequence --addr :8080",
	Short: "A web app that downloads and resizes images in sequential go routines",
	Run: func(cmd *cobra.Command, args []string) {
		sequence.Run(globalConfig)
	},
}
