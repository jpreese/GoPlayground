package main

import (
	"fmt"
)

func main() {

	var choice int

	for {

		fmt.Print("your choice: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("your number was 1")
		} else if choice == 2 {
			fmt.Println("your number was 2")
		} else if choice == 3 {
			fmt.Println("your number was 3")
		} else {
			fmt.Println("ending . . . your choice was", choice)
			break
		}
	}
}
