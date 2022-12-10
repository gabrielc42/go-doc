package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was recieved, return a value that embeds the name
	// in a greeting message
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each named people
// w/ a greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names w/ messages
	messages := make(map[string]string)
	// loop through received slice of names, calling
	// the hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved msg with the name
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		" Hi, %v. Welcome!",
		" Great to see you, %v!",
		" %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
