package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresPool(connStr string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database.")
	return pool, nil
}

func ClosePostgresPool(pool *pgxpool.Pool) {
	pool.Close()
	log.Println("Closed connection to the database.")
}
