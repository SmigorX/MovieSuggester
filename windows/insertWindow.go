package windows

import (
    //"log"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2"
)

func InsertWindow(window fyne.Window) fyne.CanvasObject {
    button := widget.NewButton("Click Me", func() {
       window.SetContent(MainWindow(window)) 
    })
    return button
}
