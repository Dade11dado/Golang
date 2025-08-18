package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	//If w name ws recived
	//Returns a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos return a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	//A map to associate names with messages
	messages := make(map[string]string)
	//Loop through the recived slice of names, calling
	//the hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//In the map associate the retrived message with the name
		messages[name] = message
	}
	return messages, nil
}

// Random format returns one of a set of greeting messages. The returned
// message on selected at random
func randomFormat() string {
	//A slice of message format
	formats := []string{
		"Hi, %v. Welcome!\n",
		"Great to see yo %v!",
		"Hail, %v! Well met!",
	}

	//return a random selected message format by specifying
	//a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
