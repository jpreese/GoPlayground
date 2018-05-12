package main

import "fmt"
import "strings"

type PasswordRule interface {
	Check(password string) bool
}

type MinimumLengthRule struct{}

func (l MinimumLengthRule) Check(password string) bool {
	const MINIMUM_PASSWORD_LENGTH = 8
	if len(password) > MINIMUM_PASSWORD_LENGTH {
		return true
	}

	return false
}

type CapitalLetterRule struct{}

func (l CapitalLetterRule) Check(password string) bool {
	const NUMBERS = "123456789"
	if strings.ContainsAny(password, NUMBERS) {
		return true
	}

	return false
}

func main() {

	input := "hello1"
	rules := []PasswordRule{MinimumLengthRule{}, CapitalLetterRule{}}

	for _, rule := range rules {
		fmt.Println(rule.Check(input))
	}
}
