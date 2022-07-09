//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

const AGE = 1000

func main() {
	switch age := AGE; {
	case age == 0:
		fmt.Println("It is a new born child!")
	case age >= 1 && age <= 3:
		fmt.Println("It is a toddler.")
	case age >= 4 && age <= 12:
		fmt.Print("It's a child")
	case age >= 13 && age <= 17:
		fmt.Println("It is a teenager.")
	case age >= 18:
		fmt.Println("It is an adult")
	default:
		fmt.Println("Age is not a number.")
	}
}
