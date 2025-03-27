package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	taxRate = getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	message := fmt.Sprintf("Ebt: %.2f, profit %.2f, ratio %.2f", ebt, profit, ratio)
	fmt.Println(message)
}

func calculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(text string) float64 {
	fmt.Print(text)
	var input float64
	fmt.Scan(&input)
	return input
}
