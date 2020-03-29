package configpi

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSSH struct {
	mock.Mock
}

func (ms *mockSSH) Execute(command string) (stdout []byte, stderr []byte, err error) {
	args := ms.Called(command)
	return args.Get(0).([]byte), args.Get(1).([]byte), args.Error(2)
}

func TestServer_Configure(t *testing.T) {
	t.Run("All commands to configure server should be executed", func(t *testing.T) {
		var mockSSH mockSSH

		mockSSH.On("Execute", mock.AnythingOfType("string")).Return([]byte("test"), []byte("test"), nil)

		server := Server{sshExec: &mockSSH, log: logrus.New()}

		err := server.Configure()

		if err != nil {
			t.Errorf("Got error %v, wanted nil", err)
		}

		mockSSH.AssertCalled(t, "Execute", aptGetUpdate)
		mockSSH.AssertCalled(t, "Execute", aptGetUpgrade)
		mockSSH.AssertCalled(t, "Execute", aptGetInstallCurl)
		mockSSH.AssertCalled(t, "Execute", installK3sServer)

		mockCalls := mockSSH.Calls

		for _, mockCall := range mockCalls {
			if mockCall.Method != "Execute" {
				t.Error("Only method 'Execute' should be called")
			}
		}

		orderArgs := []string{aptGetUpdate, aptGetUpgrade, aptGetInstallCurl, installK3sServer}

		for i, orderArg := range orderArgs {
			if arg, ok := mockCalls[i].Arguments.Get(0).(string); !ok || arg != orderArg {
				t.Errorf("First call should be with arg: %s", orderArg)
			}
		}

		if len(mockCalls) > 4 {
			t.Errorf("Execute method should not have been called more then 4 times")
		}
	})
}
