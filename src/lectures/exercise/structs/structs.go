//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length  float32
	breadth float32
}

func area(rectangle Rectangle) float32 {
	return rectangle.length * rectangle.breadth
}

func perimeter(rectangle Rectangle) float32 {
	return 2 * (rectangle.breadth + rectangle.length)
}

func main() {
	// This is a shorthand all values are required.
	r1 := Rectangle{12.5, 56}
	fmt.Println("Rectangle one", r1)

	// using var
	var (
		r2 = Rectangle{11.56, 12}
		r3 = Rectangle{}
	)

	fmt.Println(r2, r3)

	fmt.Println("Area of R1", area(r1), " and perimeter is", perimeter(r1))
	fmt.Println("Area of R1", area(r2), " and perimeter is", perimeter(r2))
	fmt.Println("Area of R1", area(r3), " and perimeter is", perimeter(r3))
}
