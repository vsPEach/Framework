package sql

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vsPEach/Framework/internal/entity"
	"github.com/vsPEach/Framework/pkg/utils"
	"log"
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

func (s *Storage) Find(ctx context.Context, model Model) (Model, error) {
	switch model.(type) {
	case entity.Article:
		var obj entity.Article
		err := s.connection.GetContext(ctx, &obj,
			fmt.Sprintf("select * from %s where id='%s'", model.GetTableName(), model.GetID()))
		return obj, err
	case entity.Comment:
		var obj entity.Comment
		err := s.connection.GetContext(ctx, &obj,
			fmt.Sprintf("select * from %s where id='%s'", model.GetTableName(), model.GetID()))
		return obj, err
	}
	return nil, errors.New("not supported")
}

func (s *Storage) FindAll(ctx context.Context, model Model) ([]Model, error) {
	var result []Model
	switch model.(type) {
	case entity.Article:
		var obj []entity.Article
		err := s.connection.SelectContext(ctx, &obj, fmt.Sprintf("select * from %s", model.GetTableName()))
		for _, article := range obj {
			result = append(result, article)
		}
		return result, err
	case entity.Comment:
		var obj []entity.Comment
		err := s.connection.SelectContext(ctx, &obj, fmt.Sprintf("select * from %s", model.GetTableName()))
		for _, article := range obj {
			result = append(result, article)
		}
		log.Println("DB GET:", result)
		return result, err
	}
	return result, errors.New("type not supported")
}

func (s *Storage) Update(ctx context.Context, model Model) error {
	query, args := utils.QueryUpdateBuilder(model)
	log.Println(query, args)
	_, err := s.connection.ExecContext(ctx, query, args...)
	return err

}

func (s *Storage) Delete(ctx context.Context, model Model) error {
	_, err := s.connection.ExecContext(
		ctx,
		fmt.Sprintf("delete from %s where id='%s'",
			model.GetTableName(),
			model.GetID(),
		))
	return err

}

func NewStorage() (*Storage, error) {
	connect, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=framework sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Storage{connection: connect}, nil
}
