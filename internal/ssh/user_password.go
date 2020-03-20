package ssh

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// UserPassword holds SSH configuration to use user and password on SSH commands
type UserPassword struct {
	host     string
	port     uint
	user     string
	password string

	log *logrus.Logger
}

// NewUserPassword creates a Executor for user and password
func NewUserPassword(host string, port uint, user string, password string, log *logrus.Logger) Executor {
	return &UserPassword{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		log:      log,
	}
}

func (up *UserPassword) Execute(command string) ([]byte, []byte, error) {
	if len(command) == 0 {
		return nil, nil, errors.New("command should not be empty")
	}

	return []byte{}, []byte{}, nil
}
