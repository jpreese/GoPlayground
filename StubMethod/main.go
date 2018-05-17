package main

import "fmt"

type DoItInterface interface {
	Do() string
}

type Thing struct{}

func (t Thing) Do() string {
	return "Real"
}

func main() {
	fmt.Println(Action(Thing{}))
}

func Action(d DoItInterface) string {
	return d.Do()
}
