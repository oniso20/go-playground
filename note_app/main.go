package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"onis-emem.com/go/note-app/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed this time.")
		return
	}

	fmt.Println("Your note has been saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)
	// var value string
	// fmt.Scanln(&value) // for single values only

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
