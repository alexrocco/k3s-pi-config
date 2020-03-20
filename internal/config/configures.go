package config

// Configures configure any setup, just a simple interface for generic configuration
type Configures interface {
	// Configure configures something and returns a error
	Configure() error
}