package greetings

import (
	"fmt"

	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name is given, return an empty string with a valid error message
    if name == "" {
        return "", errors.New("empty name")
    } 

	// Otherwise return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
