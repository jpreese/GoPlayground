package main

import "fmt"

func main() {

	input := "hello1"
	rules := []PasswordRule{MinimumLengthRule{}, ContainsNumberRule{}}

	for _, rule := range rules {
		fmt.Println(rule.Check(input))
	}
}
