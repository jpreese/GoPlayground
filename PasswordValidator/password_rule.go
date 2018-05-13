package main

type PasswordRule interface {
	Check(password string) bool
}
