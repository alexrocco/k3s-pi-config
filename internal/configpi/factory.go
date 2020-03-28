package configpi

import "github.com/sirupsen/logrus"

// Factory creates a Configuration depending on nodeType
type Factory interface {
	// Configuration creates a Configuration depending on nodeType
	Configuration(nodeType string, log *logrus.Logger) Configuration
}

// NewFactory creates a Factory for Configuration
func NewFactory() Factory {
	return &factory{}
}

type factory struct {
}

func (f *factory) Configuration(nodeType string, log *logrus.Logger) Configuration {
	switch nodeType {
	case "server":
		return NewServer(log)
	case "agent":
		return NewAgent()
	default:
		return nil
	}
}
