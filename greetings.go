package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	//If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("pls input your name first")
	}
	//Return a greeting that embeds the name in a message.
	// var message string
	// message = fmt.Sprintf("Hi , %v. Welcome!", name)
	message := fmt.Sprintf("Hi , %v. Welcome!", name)
	return message, nil
}
