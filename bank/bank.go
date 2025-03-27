package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"pl.mbednarski/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error in balance file", err)
	}

	// wantsCheckBalance := choice == 1
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("The deposit amount should be greater than zero.")
				continue
			}
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else if choice == 3 {
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("You don't have enough money. Current balance is: ", accountBalance)
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			}
		} else if choice == 4 {
			fmt.Println("Logout.")
			//Break the loop
			break
		} else {
			fmt.Println("Option is not available:", choice)
		}

	}

}
