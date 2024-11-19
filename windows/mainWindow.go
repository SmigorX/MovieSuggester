package windows

import (
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func (w *Window)mainWindow() *fyne.Container {
    test := []string{}
    data, err := w.db.ReturnMoviesByTags(test) 

    if err != nil {panic(1)}

    list := widget.NewList(
        func() int {
            return len(data)
        },
        func() fyne.CanvasObject {
            return widget.NewLabel("template")
        },
        func(i widget.ListItemID, o fyne.CanvasObject) {
            o.(*widget.Label).SetText(
                fmt.Sprintf(
                "%s, %d, %s, %v, %t",
                data[i].Name, 
                data[i].Year, 
                data[i].Director, 
                data[i].Tags, 
                data[i].Watched),
                )
        },
    )

    buttonAddWindow := widget.NewButton("Add new movie", func() {
       w.SetInsertWindow()
    })
   
    buttonExportsWindow := widget.NewButton("Exports", func() {
        w.SetExportsWindow()
    })

    buttonList := container.NewHBox(buttonAddWindow, buttonExportsWindow)

    content := container.NewBorder(
        nil,        // Top
        buttonList,     // Bottom
        nil,        // Left
        nil,        // Right
        list,       // Center
    )

    return content
}

