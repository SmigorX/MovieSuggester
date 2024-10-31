package windows

import (
    "log"

    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2"
)

func Form() fyne.Widget {
    testEntry := widget.NewEntry()
    testLongEntry := widget.NewMultiLineEntry()

    form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: testEntry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", testEntry.Text)
			log.Println("multiline:", testLongEntry.Text)
		},
	}

    return form
}

func (w *Window)exportWindow() *fyne.Container {
    button := widget.NewButton("return", func() {
       w.SetMainWindow()
    })

    container := container.NewPadded(button)
    return container
}
