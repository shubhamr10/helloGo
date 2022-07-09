package main

import "fmt"

func main() {
	var myName = "Shubham Rauniyar"
	fmt.Println("My name is ", myName)

	var name string = "Go lang"
	fmt.Println("Your name is ", name)

	userName := "shubhamr10"
	fmt.Println("My username is", userName)

	var sum int
	fmt.Println("Integer variable has default value of", sum)

	part1, other := 1, 5
	fmt.Println("Part One is", part1, "Other is", other)

	part2, other := 2, 0
	fmt.Println("Part Two is", part2, "Other is", other)

	sum = part1 + part2
	fmt.Println("Sum of Part one and part 2 is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Beginners"
	)
	fmt.Println("Lesson name is ", lessonName, " and lessson type is", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
