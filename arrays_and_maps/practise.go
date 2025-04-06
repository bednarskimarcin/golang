package main

import "fmt"

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

type Product struct {
	id    int
	title string
	price float64
}

func excercise() {
	hobbies := []string{"Guitar", "Programming", "Cycling"}
	fmt.Println("hobbies", hobbies)

	hobbiesSlice := hobbies[0:1]
	fmt.Println("hobbies slice 1", hobbiesSlice)
	hobbiesSlice = hobbiesSlice[0:2]
	fmt.Println("hobbies slice 2", hobbiesSlice)

	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Println("hobbies slice 3", hobbiesSlice)

	//7
	products := []Product{
		Product{
			id:    101,
			title: "Mac Book Pro 6",
			price: 15999.88,
		},
		Product{
			id:    102,
			title: "Lenovo ThinkPad P1",
			price: 12000.50,
		},
	}

	fmt.Println(products)

	products = append(products, Product{
		id:    102,
		title: "Dell Legion",
		price: 6500.00,
	})

	fmt.Println(products)
}
