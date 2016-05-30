package goliv

import (
	"fmt"
	"net/http"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/websocket"
)

func NewServer(opts *Options) *Server {
	if opts.Secure {
		opts.Host = "https://" + opts.Host
	} else {
		opts.Host = "http://" + opts.Host
	}

	opts.Host += ":" + opts.Port

	return &Server{
		opts: opts,
	}
}

type ServerStarter interface {
	Start() error
}

type Server struct {
	opts *Options
}

func (s *Server) Start() error {
	fmt.Println("Starting server...")

	cw := ClientWrapper{}

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(func(conn *websocket.Conn) {
		cw.Connected(conn)
	}))

	fmt.Printf("Server up: %s\n", s.opts.Port)

	open.Run(s.opts.Host)

	return http.ListenAndServe(":"+s.opts.Port, nil)
}
