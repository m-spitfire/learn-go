package greetings

import "fmt"

// Hello returns a greeting for a person
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}