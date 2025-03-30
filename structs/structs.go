package main

import (
	"fmt"

	"pl.mbednarski/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//instantiate user variable
	var appUser *user.User
	appUser, err := user.NewUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Print(err)
		return
	}

	//outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func outputUserDetails(user *user.User) {
	//Both notations work: (*user) - dereference and without it - shorthand made by Go.
	// fmt.Println((*user).firstName, user.lastName, user.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
