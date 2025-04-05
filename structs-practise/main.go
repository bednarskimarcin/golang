package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pl.mbednarski/note/note"
	"pl.mbednarski/note/todo"
)

// Naming convention: if an interface has only one method, its name should be
// the name of that method + "er"
// The interface does not have to be explicitly connected to the structs.
// The fact that the struct has the Save() method is enough and will work
type saver interface {
	Save() error
}

// An embeded interface. Contains other inreface(s) and may contain new methods.
type outputable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	text := getUserInput("TODO text")
	userNote, err := note.New(title, content)
	todo, err := todo.New(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo)

}

// Accept any type method. "any" keyword may also be used
func printSomething(value interface{}) {
	//Checking type version 1. "Is value the int type?". OK is boolean answer
	typedValue, ok := value.(int)
	if ok {
		typedValue = typedValue + 1 //We know that "int" is present here
	}

	//Version 2. Switching type. Like Java "instanceof" operator
	switch value.(type) {
	case int:
		fmt.Println("Integer ", value)
	case string:
		fmt.Println("String ", value)
	case float64:
		fmt.Println("Float64 ", value)
	}
	fmt.Print(value)
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving data failed")
		return err
	}

	fmt.Println("Saving data succeeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(promptText string) string {
	fmt.Printf("%v ", promptText)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') //Delimiter sign - terminates reading. Single qoutes "rune"

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n") //remove final linebreak
	text = strings.TrimSuffix(text, "\r")
	return text
}

// Generic type. "any" can also be used instead of "type1|...|typeN"
// But than the + operator won't work without cast.
// More generics can also be used add[T int, K float64] (a T, b K) T, K
func add[T int | float64 | string](a, b T) T {
	return a + b
}

func checkGeneric() {
	result := add(1, 2)
	fmt.Print("Generic function result ", result)
}
