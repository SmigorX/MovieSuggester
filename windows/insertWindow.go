package windows

import (
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2"
    "github.com/SmigorX/MovieSuggester/db"
    "strings"
    "strconv"
    "fmt"
)

type formFields struct {
    nameEntry *widget.Entry
    directorEntry *widget.Entry
    yearEntry *widget.Entry
    errorLabel *widget.Label
    tagsEntry *widget.Entry
    watchedEntry *widget.Check
}

func createFormFields() formFields {
    nameEntry := widget.NewEntry()
    nameEntry.SetPlaceHolder("Name")

    directorEntry := widget.NewEntry()
    directorEntry.SetPlaceHolder("Director")

    yearEntry := widget.NewEntry()
    yearEntry.SetPlaceHolder("Year")

    errorLabel := widget.NewLabel("")
    errorLabel.Hide()
    errorLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true, Monospace: true}

    tagsEntry := widget.NewEntry()
    tagsEntry.SetPlaceHolder("Tags, split by \",\"")

    watchedEntry := widget.NewCheck("Watched", func(value bool) {})

    fieldStruct := formFields{nameEntry, directorEntry, yearEntry, errorLabel, tagsEntry, watchedEntry}

    return fieldStruct
}

func parseFields(fields formFields) (int, string) {
    errors := ""

    year, err := strconv.Atoi(fields.yearEntry.Text)
    
    if err != nil {
        errors += fmt.Sprintln("Invalid year. Please enter a valid number.")
    }

    if fields.nameEntry.Text == "" {
        errors += fmt.Sprintln("Movie name cannot be empty.")
    }

    return year, errors
}

func (w *Window) fyneForm() fyne.CanvasObject {
    fields := createFormFields()

    formContainer := container.NewVBox()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Name", Widget: fields.nameEntry},
            {Text: "Director", Widget: fields.directorEntry},
            {Text: "Year", Widget: fields.yearEntry},
            {Text: "Tags", Widget: fields.tagsEntry},
            {Text: "Watched", Widget: fields.watchedEntry},
        },
        OnSubmit: func() {
            year, err := parseFields(fields)

            if err != "" {
                fields.errorLabel.SetText(err)
                fields.errorLabel.Show()
                formContainer.Refresh()
                return
            }

            fields.errorLabel.Hide()
            formContainer.Refresh()

            tags := strings.Split(fields.tagsEntry.Text, ",")
            for i, tag := range tags {
                tags[i] = strings.TrimSpace(tag)
            }
            
            movie := db.NewMovie(fields.nameEntry.Text, fields.directorEntry.Text, year, tags, fields.watchedEntry.Checked)

            w.db.AddMovie(*movie)
            fmt.Println("Movie added")
        },
    }

    // Add form and error label to the container
    formContainer.Add(fields.errorLabel)
    formContainer.Add(form)

    return formContainer
}


func (w *Window)insertWindow() *fyne.Container { 
    button := widget.NewButton("Back", func() {
       w.SetMainWindow()
    })

    form := w.fyneForm()

    content := container.NewBorder(
        nil,        // Top
        button,     // Bottom
        nil,        // Left
        nil,        // Right
        form,       // Center
    )

    return content
}
