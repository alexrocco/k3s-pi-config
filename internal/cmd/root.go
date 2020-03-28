package cmd

import (
	"fmt"
	"github.com/alexrocco/k3s-pi-config/internal/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
)

const (
	rootMsg           = "No command specified, please use the --help flag to list all the commands"
	flagNotDefinedMsg = "Error marking flags '%s' as required: "
)

func NewRoot() Commander {
	customLog := logrus.New()
	customLog.Formatter = &log.CustomFormatter{Command: "root"}

	flags := flags{}

	return &root{log: customLog, flags: flags}
}

// NewRootTest creates a config command with a custom output to be used on unit tests
func NewRootTest(out io.Writer) Commander {
	customLog := logrus.New()
	customLog.Out = out
	customLog.Formatter = &log.CustomFormatter{Command: "config"}

	flags := flags{}

	return &root{log: customLog, flags: flags}
}

type root struct {
	flags flags
	log   *logrus.Logger
}

func (r *root) Command() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "k3s-pi-config",
		Short: "k3s configuration tool for Raspberry Pi devices",
		Long:  "An easy way to configure and deploy Kubernetes clusters using k3s (https://k3s.io/) on Raspberry Pi devices.",
		Run: func(cmd *cobra.Command, args []string) {
			r.log.Info(rootMsg)
		},
	}

	// Flags
	rootCmd.PersistentFlags().StringVarP(&r.flags.host, "host", "H", "", "host name")
	rootCmd.PersistentFlags().UintVarP(&r.flags.port, "port", "P", 0, "port")
	rootCmd.PersistentFlags().StringVarP(&r.flags.user, "user", "u", "", "username")
	rootCmd.PersistentFlags().StringVarP(&r.flags.password, "password", "p", "", "password")

	// Set all flags as required
	for _, flag := range []string{"host", "port", "user", "password"} {
		err := rootCmd.MarkPersistentFlagRequired(flag)
		if err != nil {
			r.log.Errorf(fmt.Sprintf(flagNotDefinedMsg, flag), err)
			return nil
		}
	}

	// Add commands to the root command
	r.addCommands(rootCmd)

	return rootCmd
}

func (r *root) addCommands(cmd *cobra.Command) {
	// Config
	configCmd := NewConfig(&r.flags)
	cmd.AddCommand(configCmd.Command())
}
