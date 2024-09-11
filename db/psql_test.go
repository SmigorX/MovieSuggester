package db

import (
    "log"
    "os"
    "testing"
)

func TestInsertRetrieve(t *testing.T) {
    log.Printf("Starting tests")
    dbAddress := os.Getenv("MoviePSQL")
    
    if dbAddress == "" {
        t.Fatal("No db address env var")
    }

    dbInstance := New(dbAddress)

    Tags := []string{"Tag1", "Tag2"}

    TestMovie := Movie{
        Name:     "testmovie", // Assuming Name is exported
        Director: "testdirector",
        Year:     1234,
        Tags:     Tags,       // Assuming Tags is exported
        Watched:  true,
    }

    err := dbInstance.AddMovie(TestMovie)
    if err != nil {
        t.Errorf("Adding movie FAILED: %v", err)
    } else {
        t.Log("Adding movie PASSED")
    }

    movies, err := dbInstance.ReturnMoviesByTags(Tags)
    if err != nil {
        t.Errorf("Retrieving movie FAILED: %v", err)
    } else if len(movies) == 0 {
        t.Log("Retrieving movie FAILED, no movie found")
    } else {
        t.Log("Retrieving movie PASSED")
    }
}
