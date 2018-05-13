package main

import "testing"

func TestMinimumLengthRule_EmptyString_ReturnsFalse(t *testing.T) {
	result := MinimumLengthRule{}.Check("")

	if result {
		t.Error("Empty string returned", result, "should have returned", false)
	}
}

func TestMinimumLengthRule_ValidStringLength_ReturnsTrue(t *testing.T) {
	result := MinimumLengthRule{}.Check("aaaaaaaaa")

	if result == false {
		t.Error("Valid string length returned", result, "should have returned", true)
	}
}
