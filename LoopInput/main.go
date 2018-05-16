package main

import (
	"fmt"
)

func main() {

	var choice int

	for {

		fmt.Print("your choice: ")
		fmt.Scanln(&choice)

		response, e := ChoiceResolver(choice)

		if e == nil {
			fmt.Println(response)
		} else {
			fmt.Println(e)
		}
	}
}
