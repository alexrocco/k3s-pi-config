package cmd

import (
	"github.com/alexrocco/k3s-pi-config/internal/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const (
	nodeNotDefinedMsg = "node not defined, please use 'server' or 'agent'"
	wrongNodeMsg = "Wrong node defined, please use 'server' or 'agent'"
)

// NewConfig creates the config command
func NewConfig() Commander {
	customLog := logrus.New()
	customLog.Formatter = &log.CustomFormatter{Command: "config"}
	return &config{log: customLog}
}

// NewConfigTest creates a config command with a custom output to be used on unit tests
func NewConfigTest(out io.Writer) Commander {
	customLog := logrus.New()
	customLog.Out = out
	customLog.Formatter = &log.CustomFormatter{Command: "config"}
	return &config{log: customLog}
}

type config struct {
	log        *logrus.Logger
	nodeType   string
	isUnitTest bool
}

func (c *config) Command() *cobra.Command {
	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure Raspberry Pi with k3s",
		Long:  "Install k3s as a service and configure iptables to allow the communication between node and servers",
		Run:   c.run,
	}

	// Flag --node
	configCmd.Flags().StringVarP(&c.nodeType, "node", "d", "", "Node type, server or agent")
	// Flag --unit-test (this should only be used on unit tests since it disable os.Exit calls)
	configCmd.Flags().BoolVarP(&c.isUnitTest, "unit-test", "t", false, "Unit test flag, only used for development.")

	return configCmd
}

func (c *config) run(cmd *cobra.Command, args []string) {
	c.log.Infof("Init configuration for node type %s", c.nodeType)

	if len(c.nodeType) == 0 {
		c.log.Error(nodeNotDefinedMsg)
		if c.isUnitTest {
			return
		} else {
			os.Exit(1)
		}
	}

	switch c.nodeType {
	case "server":
	case "agent":
	default:
		c.log.Error(wrongNodeMsg)
		if c.isUnitTest {
			return
		} else {
			os.Exit(1)
		}
	}
}
