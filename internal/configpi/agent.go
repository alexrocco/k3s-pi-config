package configpi

func NewAgent() Configuration {
	return &Agent{}
}

type Agent struct {
}

func (s *Agent) Configure(host string, port uint, user, password string) error {
	return nil
}
