package goliv

import "fmt"

func NewServer(opts *Options) *Server {
	return &Server{
		opts: opts,
	}
}

type Server struct {
	opts *Options
}

func (s *Server) Start() {
	fmt.Println("starting server")
}
