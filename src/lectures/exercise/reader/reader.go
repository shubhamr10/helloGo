//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]string, 0)
	numOfCommands := 0

	for scanner.Scan() {
		inputText := scanner.Text()
		fmt.Println(inputText)
		numOfCommands += 1
		if inputText == "hello" || inputText == "bye" {
			fmt.Printf("Message is :%v", inputText)
		}
		if inputText == "Q" || inputText == "q" {
			lines = append(lines, inputText)
			break
		}
		if inputText == "" || inputText == "\n" {
			continue
		}
		lines = append(lines, inputText)
	}

	for scanner.Err() != nil {
		fmt.Println("Error while reading", scanner.Err())
	}

	fmt.Printf("Line counts %v \n", len(lines))
	fmt.Printf("Commands counts %v \n", numOfCommands)

}
