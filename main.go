package main

import (
	"context"
	"log"
	"os"

	"fyne.io/fyne/v2/app"

	"github.com/SmigorX/MovieSuggester/windows"

	"github.com/SmigorX/MovieSuggester/db"
)    


func database() db.DB {
    db_address := os.Getenv("MoviePSQL")
    
    if (db_address == "") {log.Fatal("No db address env var")}

    database := db.New(db_address)

    return database
}

func main() {
    db := database() 
    if db.Pool!=nil {}

    var myApp = app.New()    
    var myWindow = myApp.NewWindow("Movie Suggester")

    myWindow.SetContent(windows.MainWindow(myWindow))
    myWindow.ShowAndRun()
}

