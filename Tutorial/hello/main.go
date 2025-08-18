package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//Set the properties of the predefined Logger, including
	//the log entry prefix and a flag to disable printing
	//the time, source and file, and a line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"Davide", "Maurizio", "Luisa"}

	//request greeting messages for the names
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	//If no error returned print the returned map
	fmt.Println(messages)

	//Request greeting message
	message, err := greetings.Hello("Galdys")
	//If an error was returned, print it to the console and
	//exit the program.
	if err != nil {
		log.Fatal(err)
	}

	//If no error was returned, print the returned message
	fmt.Println(message)
}
