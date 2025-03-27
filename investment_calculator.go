package main

import (
	"fmt"
	"math"
)

// Visible globally
const inflationRate = 2.5

func main() {
	//Many variables same type in the same line
	// var investmentAmount float64 = 1000 //Either with direct type
	//var investmentAmount, years float64 = 1000, 10
	//investmentAmount, years := 1000.0, 10.0

	//Many variables different types in the same line
	//var investmentAmount, years = 1000.0, "10"

	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Please give the investment amount: ")
	fmt.Scan(&investmentAmount) //Read from standard input - single words, numbers only

	fmt.Print("Investment horizon [years]: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	message := fmt.Sprintf("Future value %.2f real value %.2f", futureValue, futureRealValue)
	outputText(message)

	multiLineMessage := `This is the 
		multiline
		message`
	outputText(multiLineMessage)
}

// Not very useful, but just for the example
func outputText(text string) {
	fmt.Println(text)
}

func calculateFactor(rate, years float64) float64 {
	result := math.Pow(1+rate/100, years)
	return result
}

func calculateFutureValues(investmentAmount, rate, years float64) (float64, float64) {
	fv := investmentAmount * calculateFactor(rate, years)
	frv := fv / calculateFactor(inflationRate, years)
	return fv, frv
}

func calculateFutureValuesAlternativeReturn(investmentAmount, rate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * calculateFactor(rate, years)
	frv = fv / calculateFactor(inflationRate, years)
	return //or you can also use: return fv, frv
}
