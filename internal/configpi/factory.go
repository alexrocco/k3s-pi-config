package configpi

// Factory creates a Configuration depending on nodeType
type Factory interface {
	// Configuration creates a Configuration depending on nodeType
	Configuration(nodeType string) Configuration
}

// NewFactory creates a Factory for Configuration
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
