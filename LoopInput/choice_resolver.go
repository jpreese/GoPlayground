package main

import (
	"errors"
)

func ChoiceResolver(choice int) (string, error) {

	var response string

	if choice == 1 {
		response = "your choice was 1"
	} else if choice == 2 {
		response = "your choice was 2"
	} else if choice == 3 {
		response = "your choice was 3"
	} else {
		return "", errors.New("fine...")
	}

	return response, nil
}
