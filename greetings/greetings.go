package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for a person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message, nil
}