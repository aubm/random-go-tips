package cmd

import (
	"github.com/aubm/random-go-tips/pkg/cmd/concurrency/pool"
	"github.com/spf13/cobra"
)

var concurrencyPoolCmd = &cobra.Command{
	Use:   "pool --addr :8080",
	Short: "A web app that downloads and resizes images in a safe concurrent way",
	Run: func(cmd *cobra.Command, args []string) {
		pool.Run(globalConfig)
	},
}
