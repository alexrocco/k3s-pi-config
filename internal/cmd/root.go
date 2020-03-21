package cmd

import (
	"github.com/alexrocco/k3s-pi-config/internal/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
)

const (
	rootMsg = "No command specified, please use the --help flag to list all the commands"
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

	// Default flags
	rootCmd.Flags().StringVarP(&r.flags.host, "host", "H", "", "host name")
	rootCmd.Flags().UintVarP(&r.flags.port, "port", "P", 0, "port")
	rootCmd.Flags().StringVarP(&r.flags.user, "user", "u", "", "username")
	rootCmd.Flags().StringVarP(&r.flags.password, "password", "p", "", "password")

	err := rootCmd.MarkFlagRequired("host")
	err = rootCmd.MarkFlagRequired("port")
	err = rootCmd.MarkFlagRequired("user")
	err = rootCmd.MarkFlagRequired("password")

	if err != nil {
		r.log.Error("Error marking flags as required: ", err)
		return nil
	}

	// Add commands to the root command
	r.addCommands(rootCmd)

	return rootCmd
}

func (r *root) addCommands(cmd *cobra.Command) {
	configCmd := NewConfig()
	cmd.AddCommand(configCmd.Command())
}
