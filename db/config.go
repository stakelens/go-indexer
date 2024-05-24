package db

import (
	"context"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
)

func Setup(connString string) (*Queries, error) {
	ctx := context.Background()
	db, err := pgxpool.New(ctx, connString)

	if err != nil {
		return nil, err
	}

	return New(db), nil
}
