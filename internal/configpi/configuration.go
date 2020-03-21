package configpi

// Configuration configure any setup, just a simple interface for generic configuration
type Configuration interface {
	// Configure configures something and returns a error
	Configure(host string, port uint, user, password string) error
}