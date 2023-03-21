package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

const (
	tag = "json"
)

func Exec(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) error {

	fmt.Println(query)

	var err error
	//* Translate ? in query
	q := db.Rebind(query)
	_, err = db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

func QueryRow[T any](ctx context.Context, db *sqlx.DB, response T, query string, args ...interface{}) (T, error) {

	fmt.Println(query)

	var err error

	//* Translate ? in query
	q := db.Rebind(query)
	//* Set json tag
	db.Mapper = reflectx.NewMapperFunc(tag, strings.ToLower)
	//* Get time type
	_, isTime := reflect.ValueOf(response).Interface().(time.Time)
	//* If not struct
	if reflect.TypeOf(response).Kind() != reflect.Struct || isTime {
		err = db.QueryRowxContext(ctx, q, args...).Scan(&response)
	} else {
		err = db.QueryRowxContext(ctx, q, args...).StructScan(&response)
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	return response, err
}

func Query[T any](ctx context.Context, db *sqlx.DB, response []T, query string, args ...interface{}) ([]T, error) {

	fmt.Println(query)

	var err error
	var compared T

	//* Get time type
	_, isTime := reflect.ValueOf(compared).Interface().(time.Time)
	//* Set json tag
	db.Mapper = reflectx.NewMapperFunc(tag, strings.ToLower)
	//* Translate ? in query
	q := db.Rebind(query)
	//* Get rows data from db
	rows, err := db.QueryxContext(ctx, q, args...)
	if err != nil && err != sql.ErrNoRows {
		return response, err
	}
	defer rows.Close()
	for rows.Next() {
		var res T
		if reflect.TypeOf(res).Kind() != reflect.Struct || isTime {
			err = rows.Scan(&res)
		} else {
			err = rows.StructScan(&res)
		}
		if err != nil && err != sql.ErrNoRows {
			return response, err
		}
		response = append(response, res)
	}
	return response, err
}
