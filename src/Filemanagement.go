package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func Save(text string) {

	//reads or creates a file called Note.txt
	file, err := os.Create(findDocumentsFolder() + "/Note.txt")
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
	file, err := os.Open(findDocumentsFolder() + "/Note.txt")
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

func findDocumentsFolder() string {

	if runtime.GOOS == "windows" {
		userProfile := os.Getenv("USERPROFILE")
		documents := filepath.Join(userProfile, "OneDrive\\Documents")
		return documents
	}
	if runtime.GOOS == "android" {
		//currently not working
		return "/Documents"
	} else {
		userProfile := os.Getenv("USERPROFILE")
		documents := filepath.Join(userProfile, "Documents")
		return documents
	}

}
