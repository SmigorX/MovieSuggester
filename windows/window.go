package windows

import (
    "github.com/SmigorX/MovieSuggester/db"
    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Window struct {
    Window fyne.Window
    db     *db.DB
}

func (w *Window) SetMainWindow() {
    w.Window.SetContent(w.mainWindow())
}

func (w *Window) SetInsertWindow() {
    w.Window.SetContent(w.insertWindow())
}

func (w *Window) SetExportsWindow() {
    w.Window.SetContent(w.exportWindow())
}

func New(db *db.DB) *Window {
    myApp := app.NewWithID("Movie Suggester")
    myWindow := myApp.NewWindow("Movie Suggester")
    newWindow := Window{Window: myWindow, db: db}
    newWindow.SetMainWindow()
    return &newWindow
}
