package main

import "testing"

type FakeDoIt struct{}

func (f FakeDoIt) Do() string {
	return "Fake"
}

func TestDo(t *testing.T) {

	result := Action(FakeDoIt{})

	if result != "Fake" {
		t.Error("Expected 'Fake' as result, but was", result)
	}

}
