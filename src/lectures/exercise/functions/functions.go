//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetPerson(name string) {
	fmt.Println("Hi", name, " welcome to golang")
}

func returnMessage(msg string) string {
	return msg
}

func addNumbers(a, b, c int) int {
	return a + b + c
}

func returnANumber(a int) int {
	return a
}

func returnTwoNumber(a, b int) (int, int) {
	return a, b
}

func main() {
	greetPerson("Shubham")
	message := returnMessage("Adios")
	fmt.Println("Message was", message)

	sum := addNumbers(2, 3, 5)
	fmt.Println("Sum is", sum)

	numOne := returnANumber(5)
	fmt.Println("Return one ", numOne)

	numTwo, numThree := returnTwoNumber(2, 3)
	fmt.Println("Two numbers are", numTwo, numThree)

	sumNumbers := addNumbers(returnANumber(5), numTwo, numThree)

	fmt.Println("Sum of three nmbrs", sumNumbers)
}
