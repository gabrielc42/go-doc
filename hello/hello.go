package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Sophie")
	fmt.Println(message)
}
