package main

import (
	"fmt"
)

func main() {
	title, content := getNoteData()

	userNote, err := Note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("Note Title:", title)
	// fmt.Println("Note Content:", content)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(promptText string) string {

	print(promptText)

	return value
}
