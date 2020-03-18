package cmd

import "github.com/spf13/cobra"

// Commander interface to expose cobra commands
type Commander interface {
	// Command that will be executed
	Command() *cobra.Command
}
