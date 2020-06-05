package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/aubm/random-go-tips/pkg/config"
)

var globalConfig = config.NewWithDefaults()

var rootCmd = &cobra.Command{
	Use:   "random-go-tips",
	Short: "A set of sample small go apps created for demo purpose",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&globalConfig.WebAppAddr, "addr", "", globalConfig.WebAppAddr, "addr to bind for web apps")

	rootCmd.AddCommand(concurrencyCmd)
	rootCmd.AddCommand(cancellationCmd)
	rootCmd.AddCommand(channelsCmd)
}
