package windows

import (
    //"log"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2"
)

func (w *Window)insertWindow() *fyne.Container { 
    button := widget.NewButton("Click Me", func() {
       w.SetMainWindow()
    })

    container := container.NewPadded(button)

    return container
}
