package main

import "strings"

type ContainsNumberRule struct{}

func (n ContainsNumberRule) Check(password string) bool {
	const NUMBERS = "123456789"
	if strings.ContainsAny(password, NUMBERS) {
		return true
	}

	return false
}
