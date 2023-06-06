package sql

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vsPEach/Framework/pkg/utils"
)

type Model interface {
	GetTableName() string
	GetID() uuid.UUID
}

type Storage struct {
	connection *sqlx.DB
	dsn        string
}

func (s *Storage) Create(ctx context.Context, model Model) error {
	query, args := utils.QueryInsertBuilder(model)
	_, err := s.connection.ExecContext(ctx, query, args...)
	return err
}

func (s *Storage) FindAll(ctx context.Context, model Model) (Model, error) {
	var obj Model
	err := s.connection.SelectContext(ctx, obj, "select * from $1", model.GetID())
	return obj, err
}

func (s *Storage) Update(ctx context.Context, model Model) error {
	query, args := utils.QueryUpdateBuilder(model)
	_, err := s.connection.ExecContext(ctx, query, args...)
	return err

}

func (s *Storage) Delete(ctx context.Context, model Model) error {
	_, err := s.connection.ExecContext(
		ctx,
		fmt.Sprintf("delete from %s where id=%s",
			model.GetTableName(),
			model.GetID(),
		))
	return err

}

func NewStorage() *Storage {
	return &Storage{}
}
