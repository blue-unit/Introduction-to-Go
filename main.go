package main

import (
	"Introduction/greetings"
	"fmt"
	"log"
)

func main() {

	// Get a greeting message and print it.
	message, err := greetings.Hello("") //Accessing the function
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
