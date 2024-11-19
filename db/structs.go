package db

type Movie struct {
    Id          int
    Name        string
    Director    string
    Year        int
    Tags        []string
    Watched     bool
}

func NewMovie(name string, director string, year int, tags []string, watched bool) *Movie {
    return &Movie{
        Id: 0,
        Name: name,
        Director: director,
        Year: year,
        Tags: tags,
        Watched: watched,
    }
}
