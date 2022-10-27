package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// NOTE: when using PascalCase the function will be exported and can be used
// from within other modules. This is not the case when using camelCase.

// NOTE: Go seems to use tuples for indicating an error. Feels way better than
// throwing. Still interested to see how typing will work for go stuff, since I
// would prefer something like Either type in Haskell.
func Hello(name string) (string, error) {
	// NOTE: := is the same as writing
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!")

	// Returns error if an empty string was passed as name
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
