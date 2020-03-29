package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRoot_Command(t *testing.T) {
	t.Run("Root command shouldn't be nil", func(t *testing.T) {
		rootCmd := NewRoot()

		cmd := rootCmd.Command()

		if cmd == nil {
			t.Error("cmd should not be nil")
		}
	})

	t.Run("Root command should output default message", func(t *testing.T) {
		var output bytes.Buffer
		rootCmd := NewRootTest(&output)

		cmd := rootCmd.Command()

		_ = cmd.PersistentFlags().Set("host", "localhost")
		_ = cmd.PersistentFlags().Set("port", "22")
		_ = cmd.PersistentFlags().Set("user", "user")
		_ = cmd.PersistentFlags().Set("password", "password")

		err := cmd.Execute()
		if err != nil {
			t.Error("err should not be nil, got: ", err)
		}

		if !strings.Contains(output.String(), rootMsg) {
			t.Error("root message not found in the output")
		}
	})

	t.Run("Root command should fail when a required flag is not set", func(t *testing.T) {
		rootCmd := NewRoot()

		cmd := rootCmd.Command()

		// Missing host flag
		_ = cmd.PersistentFlags().Set("port", "22")
		_ = cmd.PersistentFlags().Set("user", "user")
		_ = cmd.PersistentFlags().Set("password", "password")

		err := cmd.Execute()
		if err == nil {
			t.Error("err should not be nil")
		}
	})
}
