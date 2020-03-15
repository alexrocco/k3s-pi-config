package cmd

import (
	"github.com/alexrocco/k3s-pi-config/internal/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

type ConfigCmd interface {
	Command() *cobra.Command
}

func NewConfigCmd() ConfigCmd {
	customLog := logrus.New()
	customLog.Formatter = &log.CustomFormatter{Command: "config"}
	return &configCmd{log: customLog}
}

type configCmd struct {
	log      *logrus.Logger
	nodeType string
}

func (cc *configCmd) Command() *cobra.Command {
	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure Raspberry Pi with k3s",
		Long:  "Install k3s as a service and configure iptables to allow the communication between node and servers",
		Run: func(cmd *cobra.Command, args []string) {
			cc.log.Infof("Test config, node %s", cc.nodeType)

			switch cc.nodeType {
			case "server":
			case "agent":
			default:
				cc.log.Error("wrong node defined, please use 'server' or 'agent'")
				os.Exit(1)
			}
		},
	}

	configCmd.Flags().StringVarP(&cc.nodeType, "node", "d", "agent", "Node type, server or agent")

	return configCmd
}
