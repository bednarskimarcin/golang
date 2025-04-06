package main

import "fmt"

/*
See the external lib go get -u github.com/elliotchance/pie
https://medium.com/@dmytro.misik/working-with-collections-in-go-8387cecdb8a4
*/

func arraysUnpackingAndAppendSlice() {
	prices := []float64{10.99, 9.99, 45.99, 20.0}
	discountPrices := []float64{1.01, 2.01, 3.03}

	prices = append(prices, discountPrices...) //Three dots operator "..."
	fmt.Println(prices)
}

//Dynamic list
func arrays2() {
	//Without providing the size of an array, the slice is created
	//and an array correlated with the slice.
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	//Append element creates new array and returns new slice
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)
	//Remove element - no built in function. Only through slicing
	updatedPrices = updatedPrices[1:]
}

//Array fixed value
func arrays1() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices[0])
	productNames[1] = "Steve"
	fmt.Println(productNames)

	//Creating a slice
	// Right included <1, 3) left not
	// First/last index can be omitted [:3] or [1:]
	//Slice is a kind of a pointer to some part of an array (a window)
	slicePrices := prices[1:3]
	fmt.Println(slicePrices)
	//Modifying an element in the slice you modify an element in an original array
	slicePrices[0] = 6.66
	fmt.Println(slicePrices)

	//Len and cap - check the documentation for explanation
	fmt.Println(len(slicePrices), cap(slicePrices))

	//You can always reslice the slice to more elements, up to array capacity, but
	//only from going to the right. Going back to the beginning indexes does not work
	slicePrices = prices[:1]
	fmt.Println(len(slicePrices), cap(slicePrices))
	slicePrices = prices[:3]
	fmt.Println(len(slicePrices), cap(slicePrices))
}
