package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pl.mbednarski/note/note"
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
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note succeeded")
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
