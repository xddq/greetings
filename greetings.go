package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
// TODO: Why can we return nil, err here and have the return type be
// 'map[string]string, error'? Is nil a subset of the possible values of map in
// go?
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// This function will be called automaticly at program startup, after global
// variable declarations.
// src: https://go.dev/doc/effective_go#init
// "Besides initializations that cannot be expressed as declarations, a common
// use of init functions is to verify or repair correctness of the program state
// before real execution begins. "
func init() {
	// Initializes random generator with "random", time based, seed -> repeated
	// executions will probably not yield repeated results.
	rand.Seed(time.Now().UnixMilli())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
