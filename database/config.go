package database

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func Setup() (*Queries, error) {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		return nil, err
	}

	if _, err = db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return New(db), nil
}
