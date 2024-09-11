package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
    Pool *pgxpool.Pool
}

func (db *DB) ReturnMoviesByTags(tags []string) {//[]Movie {
    
}

func (db *DB) SetWatched(movie Movie) {

}

func (db *DB) AddMovie(movie Movie) {

}

func (db *DB) RemoveMovie(movie Movie) {

}


