package cmd

import (
	"bytes"
	"fmt"
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
		_ = cmd.Flags().Set("unit-test", "true")

		err := cmd.Execute()
		if err != nil {
			fmt.Println(err)
			t.Error("err should not be nil")
		}

		if !strings.Contains(output.String(), rootMsg) {
			fmt.Println(output.String())
			t.Error("root message not found in the output")
		}
	})
}