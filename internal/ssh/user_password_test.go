package ssh

import (
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewUserPassword(t *testing.T) {
	executor := NewUserPassword("test", 22, "test", "test", logrus.New())

	if executor == nil {
		t.Error("Executor should be different then nil")
	}
}

func TestUserPassword_Execute(t *testing.T) {
	// Env variable used to control the mock ssh host on pipeline
	sshHost := os.Getenv("MOCK_SSH_HOST")
	if len(sshHost) == 0 {
		sshHost = "localhost"
	}

	sshPortStr := os.Getenv("MOCK_SSH_PORT")
	sshPort := uint(2222)
	if len(sshPortStr) > 0 {
		sshPort64, err := strconv.ParseUint(sshPortStr, 10, 32)
		if err != nil {
			t.Error("error converting MOCK_SSH_PORT variable to uint")
		}

		sshPort = uint(sshPort64)
	}

	type fields struct {
		host     string
		port     uint
		user     string
		password string
		log      *logrus.Logger
	}
	type args struct {
		command string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStdout []byte
		wantStderr []byte
		wantErr    bool
	}{
		{
			name: "It should fail when the command is empty",
			fields: fields{
				host:     "test",
				port:     22,
				user:     "test",
				password: "test",
				log:      logrus.New(),
			},
			args: args{
				command: "",
			},
			wantStdout: nil,
			wantStderr: nil,
			wantErr:    true,
		},
		{
			name: "It should return the text of a echo command",
			fields: fields{
				host:     sshHost,
				port:     sshPort,
				user:     "unit-test",
				password: "test",
				log:      logrus.New(),
			},
			args: args{
				command: "echo 123",
			},
			wantStdout: []byte("123\n"),
			wantStderr: []byte(""),
			wantErr:    false,
		},
		{
			name: "It should fail when executing a command not found",
			fields: fields{
				host:     sshHost,
				port:     sshPort,
				user:     "unit-test",
				password: "test",
				log:      logrus.New(),
			},
			args: args{
				command: "fake",
			},
			wantStdout: []byte(""),
			wantStderr: []byte("bash: fake: command not found\n"),
			wantErr:    true,
		},
		{
			name: "It should fail when using a invalid port",
			fields: fields{
				host:     sshHost,
				port:     1111,
				user:     "unit-test",
				password: "test",
				log:      logrus.New(),
			},
			args: args{
				command: "fake",
			},
			wantStdout: nil,
			wantStderr: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sshExec := &UserPassword{
				host:     tt.fields.host,
				port:     tt.fields.port,
				user:     tt.fields.user,
				password: tt.fields.password,
				log:      tt.fields.log,
			}
			stdout, stderr, err := sshExec.Execute(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(stdout, tt.wantStdout) {
				t.Errorf("Execute() stdout = %v, want %v", string(stdout), string(tt.wantStdout))
			}
			if !reflect.DeepEqual(stderr, tt.wantStderr) {
				t.Errorf("Execute() stderr = %v, want %v", string(stderr), string(tt.wantStderr))
			}
		})
	}
}
