package cmd

import (
	"bytes"
	"strings"
	"testing"
)

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
		configCmd := NewConfigTest(&output)

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
		configCmd := NewConfigTest(&output)

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
		configCmd := NewConfigTest(&output)

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
		configCmd := NewConfigTest(&output)

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
