package ssh

// Executor execute SSH commands
type Executor interface {
	// Execute a SSH command
	Execute(command string) (stdout []byte, stderr[]byte, err error)
}
