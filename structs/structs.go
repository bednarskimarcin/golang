package main

import (
	"fmt"

	"pl.mbednarski/structs/user"
)

type str string // custom type - alias for string

func (s str) log() { // method for custom type
	fmt.Println("Logging:", s)
}

func main() {
	var s str = "Hello" // variable of custom type
	s.log()             // calling method for custom type

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

	admin := user.NewAdmin("mb@mbednarski.pl", "123456")
	admin.OutputUserDetails()

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
