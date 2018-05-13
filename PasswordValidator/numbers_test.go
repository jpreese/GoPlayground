package main

import "testing"

func TestContainsNumberRule_PasswordDoesNotContainNumber_ReturnsFalse(t *testing.T) {
	result := ContainsNumberRule{}.Check("a")

	if result {
		t.Error("Password without number returned", result, "should have returned", false)
	}
}

func TestContainsNumberRule_PasswordContainsNumber_ReturnsTrue(t *testing.T) {
	result := ContainsNumberRule{}.Check("1")

	if result == false {
		t.Error("Password with number returned", result, "should have returned", true)
	}
}
