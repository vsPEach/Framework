package internalhttp

import "net/http"

type Server struct {
	server http.Server
}

func NewServer() *Server {
	return &Server{server: http.Server{
		Addr:    ":8085",
		Handler: nil,
	}}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Close()
}
