package utils

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strconv"
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
	var indexes []string
	for i := range columns {
		indexes = append(indexes, "$"+strconv.Itoa(i+1))
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		model.GetTableName(),
		strings.Join(columns, ", "),
		strings.Join(indexes, ", "),
	)
	return query, values
}

func QueryUpdateBuilder(model Model) (string, []interface{}) {
	columns, values := parsingEntityQuery(model)
	sets := make([]string, 0, 10)
	for i, column := range columns {
		sets = append(sets, fmt.Sprintf("%s=$"+strconv.Itoa(i+1), column))
	}
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s';", model.GetTableName(), strings.Join(sets, ", "), model.GetID())
	return query, values
}
