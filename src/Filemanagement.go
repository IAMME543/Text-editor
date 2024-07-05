package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Save(text string) {
	//reads or creates a file called Note.txt
	file, err := os.Create("Note.rtf")
	if err != nil {
		log.Print("error creating note:", err)
	}

	defer file.Close()

	_, err = file.WriteString(text)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func LoadSavedNote() string {
	file, err := os.Open("Note.rtf")
	if err != nil {
		log.Print("There was an error opening saved note")
		return ""
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	return content

}
