package windows

import (
    "fmt"
    "os"
    "strings"
    "strconv"

    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/dialog"

    "github.com/SmigorX/MovieSuggester/db"
)

//func Form() fyne.Widget {
//    testEntry := widget.NewEntry()
//    testLongEntry := widget.NewMultiLineEntry()
//
//    form := &widget.Form{
//		Items: []*widget.FormItem{ // we can specify items in the constructor
//			{Text: "Entry", Widget: testEntry}},
//		OnSubmit: func() { // optional, handle form submission
//			log.Println("Form submitted:", testEntry.Text)
//			log.Println("multiline:", testLongEntry.Text)
//		},
//	}
//
//    return form
//}

func importFromFile(fileContent string, w *Window) {
    rows := strings.Split(fileContent, "\n")

    for _, row := range rows {
        rowContent := strings.Split(row, ";")

        if len(rowContent) != 5 {
            fmt.Printf("Invalid row %v", rowContent)
        }

        year, err := strconv.Atoi(rowContent[2])

        if err != nil {
            fmt.Printf("Invalid year %v", rowContent[2])
        }

        tags := strings.Split(rowContent[3], ",")

        watched, err := strconv.ParseBool(rowContent[4])

        if err != nil {
            fmt.Printf("Invalid watched %v", rowContent[4])
        }

        movie := db.NewMovie(rowContent[0], rowContent[1], year, tags, watched)

        err = w.db.AddMovie(*movie)

        if err != nil {
            fmt.Printf("Error adding movie %v", err)
        }
    }
} 

func importButton(w *Window) *widget.Button {
    return widget.NewButton("Import", func() {
        dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
            if err != nil {
                dialog.ShowError(err, w.Window)
                return
            }

            if reader == nil {
                return
            }

            defer reader.Close()        

            file, err := os.ReadFile(reader.URI().Path())

            if err != nil {
                dialog.ShowError(err, w.Window)
                return
            }

            importFromFile(string(file), w)

        }, w.Window)
    })
}

func exportMovies(w *Window, file *os.File) {
    movies, err := w.db.ReturnMoviesByTags([]string{})

    if err != nil {
        fmt.Printf("Error exporting movies %v", err)
        return
    }

    for _, movie := range movies {
        _, err := file.WriteString(fmt.Sprintf("%s;%s;%d;%s;%t\n", 
            movie.Name, 
            movie.Director, 
            movie.Year, 
            strings.Join(movie.Tags, ","), 
            movie.Watched),
        )

        if err != nil {
            fmt.Printf("Error writing to file %v", err)
            return
        }
    }
    
}

func exportButton(w *Window) *widget.Button {
    button := widget.NewButton("Export", func() {
        dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
            if err != nil {
                dialog.ShowError(err, w.Window)
                return
            }

            if writer == nil {
                return
            }

            defer writer.Close()

            file, err := os.Create(writer.URI().Path())

            if err != nil {
                dialog.ShowError(err, w.Window)
                return
            }

            exportMovies(w, file)
        }, w.Window)
    })

    return button
}

//For both export and import
func (w *Window)exportWindow() *fyne.Container {
    returnButton := widget.NewButton("return", func() {
       w.SetMainWindow()
    })

    importButton := importButton(w)
    exportButton := exportButton(w)

    Container := container.NewVBox(importButton, exportButton, returnButton)
    Container = container.NewCenter(Container)

    return Container
}
