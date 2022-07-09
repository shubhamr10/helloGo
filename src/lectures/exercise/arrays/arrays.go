//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func totalOfItems(products [4]Product) {
	var totalCost float32 = 0.0
	var total int = 0
	for i := 0; i < len(products); i++ {
		product := products[i]
		if product.name != "" {
			totalCost += product.price
			total++
		}

	}
	fmt.Println("Total item in the cart", total)
	fmt.Println("Total cost of items in cart", totalCost)

}

func main() {
	products := [4]Product{
		{name: "toothpaste", price: 44.56},
		{name: "soap", price: 14.06},
		{name: "shampoo", price: 110.01},
	}
	totalOfItems(products)

	products[3] = Product{name: "Pepper", price: 100.00}
	totalOfItems(products)
}
