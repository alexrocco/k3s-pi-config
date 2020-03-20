package ssh

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestNewUserPassword(t *testing.T) {
	executor := NewUserPassword("test", 22, "test", "test", logrus.New())

	if executor != nil {
		t.Error("Executor should be different then nil")
	}
}

func TestUserPassword_Execute(t *testing.T) {
	type fields struct {
		host     string
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
				host:     "",
				user:     "",
				password: "",
				log:      nil,
			},
			args:       args{},
			wantStdout: nil,
			wantStderr: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sshExec := &UserPassword{
				host:     tt.fields.host,
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
				t.Errorf("Execute() stdout = %v, want %v", stdout, tt.wantStdout)
			}
			if !reflect.DeepEqual(stderr, tt.wantStderr) {
				t.Errorf("Execute() stderr = %v, want %v", stderr, tt.wantStderr)
			}
		})
	}
}
