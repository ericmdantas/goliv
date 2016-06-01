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

	OpenBrowser(s)

	if s.opts.Secure {
		return http.ListenAndServeTLS(":"+s.opts.Port, "lib/crt/server.crt", "lib/crt/server.key", nil)
	}

	return http.ListenAndServe(":"+s.opts.Port, nil)
}

func (s *Server) OpenBrowser() error {
	return open.Run(s.opts.Host)
}
