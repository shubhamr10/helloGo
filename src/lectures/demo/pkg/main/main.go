package main

import (
	"courseContent/demo/pkg/display"
	"courseContent/demo/pkg/msg"
)

func main() {
	msg.Hi()

	display.Displays("hello from display")
	msg.Exciting("an exicting message")
}
