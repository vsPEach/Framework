package sql

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vsPEach/Framework/pkg/utils"
	"log"
	"reflect"
	"strings"
)

type Models interface {
	GetTableName() string
	GetID() uuid.UUID
}

type Storage struct {
	connection *sqlx.DB
}

func (s *Storage) Create(ctx context.Context, model Models) {
	query := strings.Builder{}
	names := make([]string, 0, 10)

	vv := reflect.ValueOf(model)

	for i := 0; i < vv.NumField(); i++ {
		names = append(names, utils.ToCamelCase(vv.Type().Field(i).Name))
	}
	for i, name := range names {
		if i == 0 {
			query.WriteString(fmt.Sprintf("insert into %s(", model.GetTableName()))
		}
		query.WriteString(fmt.Sprintf("%s, ", name))
	}
	query.WriteString(");")
	log.Println(query.String())
}

func (s *Storage) FindAll(ctx context.Context, model Models) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Update(ctx context.Context, model Models) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Delete(ctx context.Context, model Models) {
	//TODO implement me
	panic("implement me")
}

func NewStorage() *Storage {
	return &Storage{}
}
