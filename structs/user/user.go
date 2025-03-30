package user

import (
	"errors"
	"fmt"
	"time"
)

// A name of the type lowercase - private type, uppercase - public
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Adding a method to the user struct.
// It's declared outside the struct body. The (user) between func and name makes the attachement.
// That part of the code is named "receiver argument"
// However - without a pointer, that methods work on a copy of the user object
func (user User) OutputUserDetails() {
	//Here internal fields of the struct are accessed
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

// For mutator method the pointer must be used.
// Otherwise the modification is made on a copy of an object
func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

// Constructor function (by convention, not built in feature)
// That can be declared also without a pointer
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName, birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
