package configpi

func NewAgent() Configuration {
	return &Agent{}
}

type Agent struct {
}

func (s *Agent) Configure() error {
	return nil
}
