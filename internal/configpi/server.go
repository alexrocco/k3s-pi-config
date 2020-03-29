package configpi

import (
	"github.com/alexrocco/k3s-pi-config/internal/ssh"
	"github.com/sirupsen/logrus"
)

const (
	aptGetUpdate      = "sudo apt-get update"
	aptGetUpgrade     = "sudo apt-get upgrade -y"
	aptGetInstallCurl = "sudo apt-get install curl -y --no-install-recommends"
	// Install k3s (https://rancher.com/docs/k3s/latest/en/installation/install-options/
	installK3sServer = `curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server" sh -s -`
)

func NewServer(input Input, log *logrus.Logger) Configuration {
	sshExec := ssh.NewUserPassword(input.Host, input.Port, input.User, input.Password, log)

	return &Server{sshExec: sshExec, log: log}
}

type Server struct {
	sshExec ssh.Executor
	log     *logrus.Logger
}

func (s *Server) Configure() error {
	commands := []string{aptGetUpdate, aptGetUpgrade, aptGetInstallCurl, installK3sServer}

	for _, cmd := range commands {
		_, _, err := s.sshExec.Execute(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}
