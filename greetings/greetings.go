package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//Global function name i.e it can be accessed from other functions
	//outside of this module.....
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
