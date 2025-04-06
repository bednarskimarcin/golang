package main

import "fmt"

func maps1() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	mapa := map[string]int{
		"first":  3123123,
		"second": 56567,
	}
	mapa["third"] = 777

	fmt.Println(mapa)
}

func makeFunctionUsage() {
	//creates the slice
	userNames := make([]string, 2, 5) //type, initial size, capacity of the array underneath the slice
	// userNames := []string{}

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	//Creating a map with make
	courseRatings := make(floatMap, 3) //initial map allocation (3 keys-val pairs memory allocated)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	//For loops
	for index, value := range userNames {
		// ...
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		// ...
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}

//Alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}
