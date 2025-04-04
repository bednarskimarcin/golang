package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Note Title:", title)
	fmt.Println("Note Content:", content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title: ")
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	content, err := getUserInput("Note content: ")
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(promptText string) (string, error) {
	var userInput string
	print(promptText)
	fmt.Scanln(&userInput)

	if userInput == "" {
		// return "", fmt.Errorf("input cannot be empty")
		return "", errors.New("input cannot be empty")
	}

	return userInput, nil
}
