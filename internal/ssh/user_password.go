package ssh

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/alexrocco/k3s-pi-config/internal/log"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
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

	config := &ssh.ClientConfig{
		User: up.user,
		Auth: []ssh.AuthMethod{ssh.Password(up.password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			//FIXME Here it should validate the host with the public key, for now it assuming every host is valid,
			// but it should somehow validate it.
			return nil
		},
	}

	addr := fmt.Sprintf("%s:%d", up.host, up.port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, nil, err
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	logInfoWriter := log.Writer{
		Logger: up.log,
		Level:  logrus.InfoLevel,
	}

	logErrorWriter := log.Writer{
		Logger: up.log,
		Level:  logrus.DebugLevel,
	}

	session.Stdout = io.MultiWriter(&stdout, &logInfoWriter)
	session.Stderr = io.MultiWriter(&stderr, &logErrorWriter)

	up.log.Infof("SSH cmd: '%s'", command)

	err = session.Run(command)
	if err != nil {
		return stdout.Bytes(), stderr.Bytes(), err
	}

	return stdout.Bytes(), stderr.Bytes(), nil
}
