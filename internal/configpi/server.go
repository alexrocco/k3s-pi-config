package configpi

import (
	"github.com/alexrocco/k3s-pi-config/internal/ssh"
	"github.com/sirupsen/logrus"
)

func NewServer(log *logrus.Logger) Configuration {
	return &Server{log: log}
}

type Server struct {
	log *logrus.Logger
}

func (s *Server) Configure(host string, port uint, user, password string) error {
	sshExec := ssh.NewUserPassword(host, port, user, password, s.log)
	_, _, err := sshExec.Execute("echo 'test'")
	if err != nil {
		return err
	}

	return nil
}
