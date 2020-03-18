package main

import (
	"fmt"
	"github.com/alexrocco/k3s-pi-config/internal/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewRoot()

	if err := rootCmd.Command().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}