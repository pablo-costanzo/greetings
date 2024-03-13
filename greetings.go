package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	emptyNameError      = "empty name"
	emptyNamesListError = "empty list of names"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(emptyNameError)
	}
	message := fmt.Sprintf(randomGreet(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New(emptyNamesListError)
	}

	greetings := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		greetings[name] = message
	}

	return greetings, nil
}

func randomGreet() string {
	formats := []string{
		"Hi %s! Welcome!",
		"Hello %s! Nice to see you!",
		"Greetings %s! Have a great day!",
	}

	index := rand.Intn(len(formats))
	return formats[index]
}
