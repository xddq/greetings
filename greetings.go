package greetings

import "fmt"

// Hello returns a greeting for the named person.
// NOTE: when using PascalCase the function will be exported and can be used
// from within other modules. This is not the case when using camelCase.
func Hello(name string) string {
	// NOTE: := is the same as writing
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!")

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
