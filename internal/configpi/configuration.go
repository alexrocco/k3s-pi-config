package configpi

// Configuration configure any setup, just a simple interface for generic Pi configuration
type Configuration interface {
	// Configure configures a Pi and returns a error
	Configure() error
}
