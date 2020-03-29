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

	// Update all the packages
	_, _, err := sshExec.Execute("sudo apt-get update")
	if err != nil {
		return err
	}

	// Upgrade all the packages
	_, _, err = sshExec.Execute("sudo apt-get upgrade -y")
	if err != nil {
		return err
	}

	// Install cURl to download k3s bash installation file
	_, _, err = sshExec.Execute("sudo apt-get install curl -y --no-install-recommends")
	if err != nil {
		return err
	}

	// Install k3s (https://rancher.com/docs/k3s/latest/en/installation/install-options/)
	_, _, err = sshExec.Execute(`curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server" sh -s -`)
	if err != nil {
		return err
	}

	return nil
}
