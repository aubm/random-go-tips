package cmd

import (
	"github.com/aubm/random-go-tips/pkg/cmd/concurrency/sequence"
	"github.com/spf13/cobra"
)

var concurrencySequenceCmd = &cobra.Command{
	Use:   "sequence --addr :8080",
	Short: "A web app that downloads and resizes images in sequence",
	Run: func(cmd *cobra.Command, args []string) {
		sequence.Run(globalConfig)
	},
}
