package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func SelectGameFolder(w fyne.Window, cueFolderLabel *widget.Label, statusLabel *widget.Label, pathoffile *string) {
	openFolderDialog := dialog.NewFolderOpen(
		func(r fyne.ListableURI, _ error) {
			//Error showing that no folder was selected
			if r == nil {
				fmt.Println("Error: No folder selected")
				return
			}

			//Changes Label values
			cueFolderLabel.SetText("OG folder: " + r.Path())
			statusLabel.SetText("Status: Folder selected")
			*pathoffile = r.Path()

			//Selects folder where the program resides
			ex, err := os.Executable()
			if err != nil {
				panic(err)
			}
			exPath := filepath.Dir(ex)

			exeFullPath := filepath.Join(exPath, "gdi-conversion-win-v1.2.0.exe")

			//CMD running and status changing at the end depends on if there's an error or not
			kom := exec.Command(exeFullPath, "-c", r.Path())
			err = kom.Run()
			if err != nil {
				statusLabel.SetText("Error: " + err.Error())
			} else {
				statusLabel.SetText("Status: Conversion finished")
			}

		}, w)

	openFolderDialog.Show()
}

func SelectGDIFolder(w fyne.Window, cueFolderLabel *widget.Label, statusLabel *widget.Label, pathoffile *string) {
	openFolderDialog := dialog.NewFolderOpen(
		func(r fyne.ListableURI, _ error) {
			//Error showing that no folder was selected
			if r == nil {
				fmt.Println("Error: No folder selected")
				return
			}

			//Changes Label values
			cueFolderLabel.SetText("OG folder: " + r.Path())
			statusLabel.SetText("Status: Folder selected")
			*pathoffile = r.Path()

			//Selects folder where the program resides
			ex, err := os.Executable()
			if err != nil {
				panic(err)
			}
			exPath := filepath.Dir(ex)

			exeFullPath := filepath.Join(exPath, "gdi-conversion-win-v1.2.0.exe")

			//CMD running and status changing at the end depends on if there's an error or not
			kom := exec.Command(exeFullPath, "-n", r.Path())
			err = kom.Run()
			if err != nil {
				statusLabel.SetText("Error: " + err.Error())
			} else {
				statusLabel.SetText("Status: Extraction finished")
			}

		}, w)

	openFolderDialog.Show()
}

func main() {
	a := app.New()
	w := a.NewWindow("GDI Conversion GUI 0.1")
	pathoffile := ""
	var buttonChoseFolder *widget.Button
	var statusLabel *widget.Label

	cueFolderLabel := widget.NewLabel("OG folder: ")
	statusLabel = widget.NewLabel("Status: Conversion have not started yet")

	check := widget.NewCheck("check if you want to extract name of GDI instead", func(b bool) {
		if b != false {
			buttonChoseFolder.SetText("Select folder with gdi and extract name")
			statusLabel.SetText("Status: Name extraction have not started yet")
		} else {
			buttonChoseFolder.SetText("Select folder with cue and convert")
			statusLabel.SetText("Status: Conversion have not started yet")
		}
	})

	buttonChoseFolder = widget.NewButton("Select folder with cue and convert", func() {
		if check.Checked {
			buttonChoseFolder.SetText("Select folder with gdi and extract name")
			SelectGDIFolder(w, cueFolderLabel, statusLabel, &pathoffile)
		} else {
			buttonChoseFolder.SetText("Select folder with cue and convert")
			SelectGameFolder(w, cueFolderLabel, statusLabel, &pathoffile)
		}

	})

	vbox := container.NewVBox(check, buttonChoseFolder, cueFolderLabel, statusLabel)

	w.SetContent(vbox)
	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()
}

//Author: Maciej Piątek (2025)
//mpdev@memeware.net

//Icon from https://www.hiclipart.com/free-transparent-background-png-clipart-qzqwo/download

//Do not modify without creditting me and letting me know by mail!
