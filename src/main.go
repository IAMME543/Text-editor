package main

/*INformaition
- app id = com.ATextEditor.ATE
*/

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	//initializes the window and application
	a := app.NewWithID("A-Text-Editor-IamMe")

	w := a.NewWindow("Text Editor")

	a.Settings().SetTheme(&myTheme{})

	windowSize := fyne.NewSize(700, 750)

	// textDisplay := widget.NewRichText()
	// textDisplay.ParseMarkdown(LoadSavedNote())

	//set up entry
	inputFeild := NewCustomEntry()
	inputFeild.MultiLine = true
	inputFeild.Text = LoadSavedNote()

	inputFeild.OnChanged = func(text string) {
		Save(text)
		// textDisplay.ParseMarkdown(LoadSavedNote())
	}

	//parentCon := container.NewVBox(textDisplay, inputFeild)

	//creates the window and launch
	w.SetContent(inputFeild)

	w.Resize(windowSize)
	w.CenterOnScreen()

	w.ShowAndRun()

}
