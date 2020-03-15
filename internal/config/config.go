package config

type Configurer interface {
	Configure(nodeType string) error

}
