package internalhttp

import (
	"context"
	"github.com/vsPEach/Framework/internal/repository/sql"
	"github.com/vsPEach/Framework/internal/server/http/routes"
	"net/http"
)

type Storage interface {
	Create(ctx context.Context, model sql.Model) error
	Update(ctx context.Context, model sql.Model) error
	FindAll(ctx context.Context, model sql.Model) ([]sql.Model, error)
	Delete(ctx context.Context, model sql.Model) error
	Find(ctx context.Context, model sql.Model) (sql.Model, error)
}

type Server struct {
	server http.Server
}

func NewServer(storage Storage) *Server {
	r := routes.NewHTTPHandler(storage)
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
