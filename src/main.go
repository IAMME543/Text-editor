package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	//initializes the window and application
	a := app.New()
	w := a.NewWindow("Text Editor")
	a.Settings().SetTheme(&myTheme{})

	windowSize := fyne.NewSize(700, 750)

	//set up entry
	inputFeild := NewCustomEntry()
	inputFeild.MultiLine = true
	inputFeild.Text = LoadSavedNote()

	inputFeild.OnChanged = func(text string) {
		SaveFile(text)
	}

	//creates the window and launch
	w.SetContent(inputFeild)

	w.Resize(windowSize)
	w.CenterOnScreen()

	w.ShowAndRun()

}

func SaveFile(text string) {
	file, err := os.Create("Note.txt")
	if err != nil {
		log.Print("Error creating file 'Note.txt'")
	}
	defer file.Close()

	_, err = file.WriteString(text)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}

func LoadSavedNote() string {
	file, err := os.Open("Note.txt")
	if err != nil {
		log.Print("There was an error opening saved note")
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	return content

}
