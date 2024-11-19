package db

import (
    "log"
    "context"

    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
    Pool *pgxpool.Pool
    Ctx  context.Context   
}

//New creates a new db connection
func New (connectionString string) *DB {
    new_db := &DB{
        Pool: NewPostgresPool(connectionString),
        Ctx: context.Background(),
    }
    return new_db
}

//ReturnMoviesByTags takes list of tags and returns a list of all matched movies
func (db *DB) ReturnMoviesByTags(tags []string) ([]Movie, error) {
    var rows pgx.Rows 
    var err error
    if len(tags) == 0 {
        rows, err = db.Pool.Query(db.Ctx, "SELECT id, name, director, year, tags, watched FROM movies")
    } else {
        rows, err = db.Pool.Query(db.Ctx, "SELECT id, name, director, year, tags, watched FROM movies WHERE tags && $1",  tags)
    }

    if (err != nil) {
        log.Printf("ReturnMoviesByTags failed %v", err)
        return nil, err
    }
    defer rows.Close()

    var movies []Movie

    for rows.Next() {
        var movie Movie

        err := rows.Scan(&movie.Id, &movie.Name, &movie.Director, &movie.Year, &movie.Tags, &movie.Watched)
    
        if err != nil {
            log.Printf("Scanning rows failed %v", err)
        } else {
            movies = append(movies, movie)
        }
    }

    return movies, nil
}

//ChangeWatched finds a row by movie's id and flips the watched column to opposite value
func (db *DB) ChangeWatched(movie Movie) error {
    _, err := db.Pool.Exec(db.Ctx, "UPDATE movies SET watched = NOT watched WHERE id = $1", movie.Id)

    if (err != nil) {
        log.Printf("ChangeWatched failed %v", err)
        return err
    }
    return nil
}

//AddMovie adds new entry to the movie db
func (db *DB) AddMovie(movie Movie) error {
    row := db.Pool.QueryRow(
        db.Ctx, 
        `INSERT INTO movies (name, director, watched, year, tags) 
        VALUES ($1, $2, $3, $4, $5)`,
        movie.Name, movie.Director, movie.Watched, movie.Year, movie.Tags,
    )


    // Check if the row was successfully returned, or if an error occurred
    err := row.Scan() // Scan the returned row (if the query was successful)
    if err != nil && err != pgx.ErrNoRows {
        log.Printf("AddMovie failed, %v", err)
        return err
    }

    // No need for err.Scan() here since it's not applicable with pgx
    return nil
}

//RemoveMovie removes a row with movie's id
func (db *DB) RemoveMovie(movie Movie) error {
    _, err := db.Pool.Exec(db.Ctx, "DELETE FROM movies WHERE id = $1", movie.Id)
    
    if err != nil {
        log.Printf("RemoveMovie failed, %v", err)
        return err
    }
    return nil
}


