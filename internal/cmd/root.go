package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k3s-pi-config",
	Short: "k3s configuration tool for Raspberry Pi devices",
	Long:  "An easy way to configure and deploy Kubernetes clusters using k3s (https://k3s.io/) on Raspberry Pi devices.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Just a test")
	},
}

// Execute runs the default command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
