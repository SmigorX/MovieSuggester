package main

import (
    "log"

	"fyne.io/fyne/v2/app"
    
    "github.com/SmigorX/MovieSuggester/windows"

    "github.com/SmigorX/MovieSuggester/db"
)    

var db_address string = "postgres://username:password@localhost:5432/mydb"

func database(db_address string) db.DB {
    pool, err := db.NewPostgresPool(db_address)

    if (err != nil) {
        log.Fatal("Could not connect to the database")
    }

    database := db.DB{Pool: pool}

    return database
}

func main() {
    db := database(db_address) 
    if db.Pool!=nil {}

    var myApp = app.New()    
    var myWindow = myApp.NewWindow("Movie Suggester")

    myWindow.SetContent(windows.MainWindow(myWindow))
    myWindow.ShowAndRun()
}

