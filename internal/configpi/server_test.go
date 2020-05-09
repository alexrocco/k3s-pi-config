package configpi

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type mockSSH struct {
	mock.Mock
}

func (ms *mockSSH) Execute(command string) (stdout []byte, stderr []byte, err error) {
	args := ms.Called(command)

	stdout = nil
	if args.Get(0) != nil {
		stdout = args.Get(0).([]byte)
	}

	stderr = nil
	if args.Get(0) != nil {
		stderr = args.Get(1).([]byte)
	}

	return stdout, stderr, args.Error(2)
}

func TestServer_Configure(t *testing.T) {
	t.Run("All commands to configure server should be executed", func(t *testing.T) {
		var mockSSH mockSSH

		mockSSH.On("Execute", mock.AnythingOfType("string")).Return([]byte("test"), []byte("test"), nil)

		server := server{sshExec: &mockSSH, log: logrus.New()}

		err := server.Configure()

		if err != nil {
			t.Errorf("Got error %v, wanted nil", err)
		}

		mockSSH.AssertCalled(t, "Execute", aptGetUpdate)
		mockSSH.AssertCalled(t, "Execute", aptGetUpgrade)
		mockSSH.AssertCalled(t, "Execute", aptGetInstallCurl)
		mockSSH.AssertCalled(t, "Execute", installK3sServer)

		for _, mockCall := range mockSSH.Calls {
			if mockCall.Method != "Execute" {
				t.Error("Only method 'Execute' should be called")
			}
		}

		orderArgs := []string{aptGetUpdate, aptGetUpgrade, aptGetInstallCurl, installK3sServer}

		for i, orderArg := range orderArgs {
			if arg, ok := mockSSH.Calls[i].Arguments.Get(0).(string); !ok || arg != orderArg {
				t.Errorf("First call should be with arg: %s", orderArg)
			}
		}

		if len(mockSSH.Calls) > 4 {
			t.Errorf("Execute method should not have been called more then 4 times")
		}
	})

	t.Run("When a command fail the next commands must not be called", func(t *testing.T) {
		var mockSSH mockSSH

		someErr := errors.New("some error")
		mockSSH.On("Execute", aptGetInstallCurl).Return(nil, nil, someErr)

		mockSSH.On("Execute", mock.AnythingOfType("string")).Return([]byte("test"), []byte("test"), nil)

		server := server{sshExec: &mockSSH, log: logrus.New()}

		err := server.Configure()

		mockSSH.AssertCalled(t, "Execute", aptGetUpdate)
		mockSSH.AssertCalled(t, "Execute", aptGetUpgrade)
		mockSSH.AssertCalled(t, "Execute", aptGetInstallCurl)
		mockSSH.AssertNotCalled(t, "Execute", installK3sServer)


		if err == nil || !reflect.DeepEqual(someErr, err) {
			t.Errorf("Got error %v, wanted %v", err, someErr)
		}

		orderArgs := []string{aptGetUpdate, aptGetUpgrade, aptGetInstallCurl}

		for i, orderArg := range orderArgs {
			if arg, ok := mockSSH.Calls[i].Arguments.Get(0).(string); !ok || arg != orderArg {
				t.Errorf("First call should be with arg: %s", orderArg)
			}
		}

		if len(mockSSH.Calls) > 3 {
			t.Errorf("Execute method should not have been called more then 3 times")
		}


	})
}
