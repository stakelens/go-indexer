package database

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func Setup(driverName, dataSourceName string) (*Queries, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	if _, err = db.ExecContext(context.Background(), ddl); err != nil {
		return nil, err
	}

	return New(db), nil
}
