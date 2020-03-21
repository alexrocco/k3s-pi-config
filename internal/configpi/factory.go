package configpi

type Factory interface {
	Configuration(nodeType string) Configuration
}

func NewFactory() Factory {
	return &factory{}
}

type factory struct {
}

func (f *factory) Configuration(nodeType string) Configuration {
	switch nodeType {
	case "server":
		return NewServer()
	case "agent":
		return NewAgent()
	default:
		return nil
	}
}
