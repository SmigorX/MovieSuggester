package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresPool(connStr string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("ERROR whilst creating a pool config, %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("ERROR whilst connecting to the pool, %v", err)
	}
    defer pool.Close()

	log.Println("Connected to the database.")
	return pool
}

