package cmd

import "github.com/spf13/cobra"

// Commander interface to expose cobra commands
type Commander interface {
	// Command that will be executed
	Command() *cobra.Command
}

// flags wraps all the default flags that will be used on all sub-commands
type flags struct {
	host     string
	port     uint
	user     string
	password string
}
