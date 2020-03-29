package cmd

import (
	"io"
	"os"

	"github.com/alexrocco/k3s-pi-config/internal/configpi"
	"github.com/alexrocco/k3s-pi-config/internal/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	nodeNotDefinedMsg = "Node not defined, please use 'server' or 'agent'"
	wrongNodeMsg      = "Wrong node defined, please use 'server' or 'agent'"
	configErrorMsg    = "Error while configuring '%s': %v "
)

// NewConfig creates the config command
func NewConfig(flags *flags) Commander {
	customLog := logrus.New()
	customLog.Formatter = &log.CustomFormatter{Command: "config"}

	configpiFactory := configpi.NewFactory()

	return &config{log: customLog, configpiFactory: configpiFactory, flags: flags}
}

// NewConfigTest creates a config command with a custom output to be used on unit tests
func NewConfigTest(out io.Writer, configpiFactory configpi.Factory) Commander {
	customLog := logrus.New()
	customLog.Out = out
	customLog.Formatter = &log.CustomFormatter{Command: "config"}


	return &config{log: customLog, configpiFactory: configpiFactory, flags: &flags{}}
}

type config struct {
	flags      *flags
	nodeType   string
	isUnitTest bool

	configpiFactory configpi.Factory
	log             *logrus.Logger
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
		}

		os.Exit(1)
	}

	config := c.configpiFactory.Configuration(c.nodeType, c.log)
	if config == nil {
		c.log.Error(wrongNodeMsg)

		if c.isUnitTest {
			return
		}

		os.Exit(1)
	}

	err := config.Configure(c.flags.host, c.flags.port, c.flags.user, c.flags.password)
	if err != nil {
		c.log.Errorf(configErrorMsg, c.nodeType, err)

		if c.isUnitTest {
			return
		}

		os.Exit(1)
	}
}
