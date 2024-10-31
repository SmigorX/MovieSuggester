package main

import (
	//"context"
	"log"
	"os"


	"github.com/SmigorX/MovieSuggester/windows"

	"github.com/SmigorX/MovieSuggester/db"
)    


func database() *db.DB {
    db_address := os.Getenv("MoviePSQL")
    
    if (db_address == "") {log.Fatal("No db address env var")}

    database := db.New(db_address)

    return database
}

func main() {
    db := database() 
    if db.Pool==nil {log.Print("No pool")}

    window := windows.New(db)
    window.Window.ShowAndRun()
}

