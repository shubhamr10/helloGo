package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 scored higher than quiz 2 ")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 2 scored higher than quiz 1 ")
	} else {
		fmt.Println("quiz1 and quiz has the same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("acceptable grades")
	} else {
		fmt.Println("instructor is ")
	}
}
