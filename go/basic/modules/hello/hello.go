package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of defined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Slice of names
	names := []string{"bread", "potato", "tomato"}

	// Request a greeting message
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// Get a greeting message and print it.
	// message := greetings.Hello("Gladys")
	fmt.Println(message)
}
