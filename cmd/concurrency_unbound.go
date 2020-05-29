package cmd

import (
	"github.com/aubm/random-go-tips/pkg/cmd/concurrency/unbound"
	"github.com/spf13/cobra"
)

var concurrencyUnboundCmd = &cobra.Command{
	Use:   "unbound --addr :8080",
	Short: "A web app that compute fibonacci numbers in an unbound concurrent way",
	Run: func(cmd *cobra.Command, args []string) {
		unbound.Run(globalConfig)
	},
}
