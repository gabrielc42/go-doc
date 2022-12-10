package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of predefined logger, including
	// log entry prefix and a flag to
	// 	disable printing time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// request greeting messages for the names
	messages, err := greetings.Hellos(names)

	// Request a greeting message
	// message, err := greetings.Hello("Gladys")

	// if an error was returned, print it to the console
	// and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print returned msg
	fmt.Println(messages)
}
