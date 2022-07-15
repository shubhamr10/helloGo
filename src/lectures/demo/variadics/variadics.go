package main

import "fmt"

func sums(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{4, 5, 6}

	all := append(a, b...)

	answer := sums(all...)
	fmt.Println("Answer", answer)

	answer = sums(1, 2, 3, 4, 4, 5, 6)
	fmt.Println("Answer", answer)

}
