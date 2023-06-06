package utils

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

const tag = "sqlx"

type Model interface {
	GetTableName() string
	GetID() uuid.UUID
}

func parsingEntityQuery(model Model) ([]string, []interface{}) {
	value := reflect.ValueOf(model)
	typ := reflect.TypeOf(model)
	var values []interface{}
	var columns []string

	for i := 0; i < value.NumField(); i++ {
		fieldVal := value.Field(i)
		fieldType := typ.Field(i)

		tagName := fieldType.Tag.Get(tag)
		if tagName != "" {

			columns = append(columns, tagName)

			values = append(values, fieldVal.Interface())
		}
	}
	return columns, values
}

func QueryInsertBuilder(model Model) (string, []interface{}) {
	columns, values := parsingEntityQuery(model)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		model.GetTableName(),
		strings.Join(columns, ", "),
		strings.TrimSuffix(strings.Repeat("?, ", len(values)), ", "))
	return query, values
}

func QueryUpdateBuilder(model Model) (string, []interface{}) {
	columns, values := parsingEntityQuery(model)
	values = append(values, model.GetID())
	sets := make([]string, 0, 10)
	for _, column := range columns {
		sets = append(sets, fmt.Sprintf("%s=?", column))
	}
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = ?;", strings.Join(sets, ", "))
	return query, values
}
