package cmd

import (
	"bytes"
	"github.com/alexrocco/k3s-pi-config/internal/configpi"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

type (
	mockFactory struct{}

	mockServer struct {
		mock.Mock
	}

	mockAgent struct {
		mock.Mock
	}
)

func (mf *mockFactory) Configuration(nodeType string) configpi.Configuration {
	switch nodeType {
	case "server":
		return &mockServer{}
	case "agent":
		return &mockAgent{}
	default:
		return nil
	}
}

func (ms *mockServer) Configure(host string, port uint, user, password string) error {
	args := ms.Called(host, port, user, password)
	return args.Error(0)
}

func (ma *mockAgent) Configure(host string, port uint, user, password string) error {
	args := ma.Called(host, port, user, password)
	return args.Error(0)
}

func TestConfig_Command(t *testing.T) {
	configCmd := NewConfig()

	cmd := configCmd.Command()

	if cmd == nil {
		t.Error("cmd should not be nil")
	}
}

func TestConfig_run(t *testing.T) {
	t.Run("Run should output when node flag is empty", func(t *testing.T) {
		var output bytes.Buffer
		configCmd := NewConfigTest(&output, &mockFactory{})

		cmd := configCmd.Command()
		_ = cmd.Flags().Set("unit-test", "true")

		// Execute will call the run method
		err := cmd.Execute()
		if err != nil {
			t.Error("err should not be nil")
		}

		if !strings.Contains(output.String(), nodeNotDefinedMsg) {
			t.Error("message node not defined not found in output")
		}
	})

	t.Run("Run should output when node value is wrong", func(t *testing.T) {
		var output bytes.Buffer
		configCmd := NewConfigTest(&output, &mockFactory{})

		cmd := configCmd.Command()
		_ = cmd.Flags().Set("unit-test", "true")
		_ = cmd.Flags().Set("node", "test")

		err := cmd.Execute()
		if err != nil {
			t.Error("err should not be nil")
		}

		if !strings.Contains(output.String(), wrongNodeMsg) {
			t.Error("message wrong node not found in output")
		}
	})

	t.Run("Run should work when the node type is 'agent'", func(t *testing.T) {
		var output bytes.Buffer
		configCmd := NewConfigTest(&output, &mockFactory{})

		cmd := configCmd.Command()
		_ = cmd.Flags().Set("unit-test", "true")
		_ = cmd.Flags().Set("node", "agent")

		err := cmd.Execute()
		if err != nil {
			t.Error("err should not be nil")
		}

		if strings.Contains(output.String(), wrongNodeMsg) || strings.Contains(output.String(), wrongNodeMsg) {
			t.Error("validations outputs shouldn't be logged")
		}
	})

	t.Run("Run should work when the node type is 'server'", func(t *testing.T) {
		var output bytes.Buffer
		configCmd := NewConfigTest(&output, &mockFactory{})

		cmd := configCmd.Command()
		_ = cmd.Flags().Set("unit-test", "true")
		_ = cmd.Flags().Set("node", "server")

		err := cmd.Execute()
		if err != nil {
			t.Error("err should not be nil")
		}

		if strings.Contains(output.String(), wrongNodeMsg) || strings.Contains(output.String(), wrongNodeMsg) {
			t.Error("validations outputs shouldn't be logged")
		}
	})
}
