package main

import "fmt"

func main() {
	age := 32 //regular variable

	var agePointer *int
	agePointer = &age //memory address of the age variable

	fmt.Println("Age: ", *agePointer) //dereferencing - get the value of the pointer (not an address)

	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult years: ", adultYears)

	editAgeToAdultYears(agePointer)
	fmt.Println("Age after mutate: ", age)
}

func getAdultYears(age *int) int {
	return *age - 18 //dereference the pointer to be able to perform arythmetics on values
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18 //directly mutate the value of the pointer
}
