package internalhttp

import (
	"github.com/vsPEach/Framework/internal/server/http/routes"
	"net/http"
)

type Server struct {
	server http.Server
}

func NewServer() *Server {
	r := routes.HTTPHandler{}
	return &Server{server: http.Server{
		Addr:    ":8085",
		Handler: r.Routes(),
	}}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Close()
}
