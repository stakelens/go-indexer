package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Setup(connString string) (*Queries, error) {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, connString)

	if err != nil {
		return nil, err
	}

	return New(db), nil
}
