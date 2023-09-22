package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If name exists
	message := fmt.Sprintf("Hello, %v. Welcome to Go", name)
	return message, nil
}