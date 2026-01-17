package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB() error {
	var err error

	DB, err = pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/shortener")
	if err != nil {
		return err
	}

	return DB.Ping(context.Background())
}
