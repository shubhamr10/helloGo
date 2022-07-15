//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Items struct {
	name string
	tag  bool
}

func tagsChange(tag *bool) {
	*tag = !*tag
}

func checkout(items []Items) {
	fmt.Println("checking out...")
	for i := 0; i < len(items); i++ {
		tagsChange(&items[i].tag)
	}
	fmt.Println(items, 'a')
}

func main() {
	items := []Items{
		{name: "Bread", tag: true},
		{name: "belt", tag: true},
		{name: "bottle", tag: true},
		{name: "tape", tag: true},
	}

	fmt.Println(items)
	tagsChange(&items[2].tag)
	fmt.Println(items)

	checkout(items)
	fmt.Println(items)

}
