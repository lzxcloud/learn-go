package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("lzx")
	fmt.Println(message)
}