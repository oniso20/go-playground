package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"onis-emem.com/go/note-app/note"
	"onis-emem.com/go/note-app/todo"
)

// Interfaces
type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

// Embedded Interfaces that works with interfaces live saver or just the method like Display()
type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()

	todoText := getUserInput("Todo Content: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Display and save todo

	err = outputData(todo)

	if err != nil {
		return
	}

	// Display and Save notes

	outputData(userNote)
}

func printSomething(value interface{}) {

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Integer:", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("Integer:", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer", value)
	// case float64:
	// 	fmt.Println("Float", value)
	// case string:
	// 	fmt.Println("String", value)
	// }

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed this time.")
		return err
	}

	fmt.Println("Your note has been saved successfully!")

	return nil
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
