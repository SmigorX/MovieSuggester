package db

type Movie struct {
    Id          int
    Name        string
    Director    string
    Year        int
    Tags        []string
    Watched     bool
}

