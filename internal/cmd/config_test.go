package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/alexrocco/k3s-pi-config/internal/configpi"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

type (
	mockFactory struct {
		mockServer *mockServer
		mockAgent  *mockAgent
	}

	mockServer struct {
		mock.Mock
	}

	mockAgent struct {
		mock.Mock
	}
)

func (mf *mockFactory) Configuration(nodeType string, log *logrus.Logger) configpi.Configuration {
	switch nodeType {
	case "server":
		return mf.mockServer
	case "agent":
		return mf.mockAgent
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
	configCmd := NewConfig(&flags{})

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
		var mockAgent mockAgent

		mockAgent.On("Configure",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("uint"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).
			Return(nil)

		configCmd := NewConfigTest(&output, &mockFactory{mockAgent: &mockAgent})

		cmd := configCmd.Command()
		_ = cmd.Flags().Set("unit-test", "true")
		_ = cmd.Flags().Set("node", "agent")

		err := cmd.Execute()
		if err != nil {
			t.Error("err should not be nil")
		}

		mockAgent.AssertNumberOfCalls(t, "Configure", 1)

		if strings.Contains(output.String(), wrongNodeMsg) {
			t.Error("validations outputs shouldn't be logged")
		}
	})

	t.Run("Run should work when the node type is 'server'", func(t *testing.T) {
		var output bytes.Buffer
		var mockServer mockServer

		mockServer.On("Configure",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("uint"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).
			Return(nil)

		configCmd := NewConfigTest(&output, &mockFactory{mockServer: &mockServer})

		cmd := configCmd.Command()
		_ = cmd.Flags().Set("unit-test", "true")
		_ = cmd.Flags().Set("node", "server")

		err := cmd.Execute()
		if err != nil {
			t.Error("err should not be nil")
		}

		mockServer.AssertNumberOfCalls(t, "Configure", 1)

		if strings.Contains(output.String(), wrongNodeMsg) {
			t.Error("validations outputs shouldn't be logged")
		}
	})

	t.Run("Run should fail when configure process fails", func(t *testing.T) {
		var output bytes.Buffer
		var mockServer mockServer

		nodeType := "server"
		mockErr := errors.New("mock error")

		mockServer.On("Configure",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("uint"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).
			Return(mockErr)

		configCmd := NewConfigTest(&output, &mockFactory{mockServer: &mockServer})

		cmd := configCmd.Command()
		_ = cmd.Flags().Set("unit-test", "true")
		_ = cmd.Flags().Set("node", nodeType)

		err := cmd.Execute()
		if err != nil {
			t.Error("err should be nil")
		}

		mockServer.AssertNumberOfCalls(t, "Configure", 1)

		if !strings.Contains(output.String(), fmt.Sprintf(configErrorMsg, nodeType, mockErr)) {
			t.Error("wrong error message")
		}
	})
}
