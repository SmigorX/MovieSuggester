package db

import (
    "log"
    "context"
	
    "github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
    Pool *pgxpool.Pool
    Ctx  context.Context   
}

func New (connectionString string) *DB {
    new_db := &DB{
        Pool: NewPostgresPool(connectionString),
        Ctx: context.Background(),
    }
    return new_db
}

func (db *DB) ReturnMoviesByTags(tags []string) ([]Movie, error) {
    rows, err := db.Pool.Query(db.Ctx, "SELECT id, name, director, year, tags, watched FROM movies")

    if (err != nil) {
        log.Printf("Querry failed %v", err)
        return nil, err
    }
    defer rows.Close()

    var movies []Movie

    for rows.Next() {
        var movie Movie
        var ignoredid int

        err := rows.Scan(ignoredid, movie.Name, movie.Director, movie.Year, movie.Tags, movie.Watched)
    
        if err != nil {
            log.Printf("Scanning rows failed %v", err)
        } else {
            movies = append(movies, movie)
        }
    }
    
    return movies, nil
}

func (db *DB) SetWatched(movie Movie) {

}

func (db *DB) AddMovie(movie Movie) error {
    err := db.Pool.QueryRow(
        db.Ctx, 
        `INSERT INTO movies (name, director, watched, year) 
        VALUES ($1, $2, $3, $4)`,
        movie.Name, movie.Director, movie.Watched, movie.Year,
    )

    if err != nil {
        log.Printf("Insert failed, %v", err)
        return err.Scan()
    }

    return nil
}

func (db *DB) RemoveMovie(movie Movie) {

}


