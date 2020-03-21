package configpi

func NewServer() Configuration {
	return &Server{}
}

type Server struct {
}

func (s *Server) Configure(host string, port uint, user, password string) error {
	return nil
}