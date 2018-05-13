package main

type MinimumLengthRule struct{}

func (l MinimumLengthRule) Check(password string) bool {
	const MINIMUM_PASSWORD_LENGTH = 8
	if len(password) > MINIMUM_PASSWORD_LENGTH {
		return true
	}

	return false
}
