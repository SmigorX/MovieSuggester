package windows

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func MainWindow(window fyne.Window) fyne.CanvasObject {
    data := []string{"a", "string", "list", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"} //DEV

    list := widget.NewList(
        func() int {
            return len(data)
        },
        func() fyne.CanvasObject {
            return widget.NewLabel("template")
        },
        func(i widget.ListItemID, o fyne.CanvasObject) {
            o.(*widget.Label).SetText(data[i])
        },
    )

    button := widget.NewButton("Click Me", func() {
       window.SetContent(InsertWindow(window)) 
    })
    
    content := container.NewBorder(
        nil,        // Top
        button,     // Bottom
        nil,        // Left
        nil,        // Right
        list,       // Center
    )

    return content
}

