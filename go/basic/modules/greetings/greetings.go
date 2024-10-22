package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return error if no name supplied
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	// num := len(message)
	return message, nil
}

// Hellos returns a map that associates each of the named people with
// a greeting message
func Hellos(names []string) (map[string]string, error) {
	// Map to associate names with messages
	messages := make(map[string]string)
	// Loop through received slice of names, calling
	// Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// Assocate retrieved message with name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns a random greeting messages. The returned
// message is selected at random
func randomFormat() string {
	// Slice of message formats
	formats := []string{
		"Hi, %v.",
		"See you %v",
	}

	// Return randomly selected message format using random index
	// for slice of formats
	return formats[rand.Intn(len(formats))]
}
