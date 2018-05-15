package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string

	for {

		fmt.Print("your choice: ")
		fmt.Scanln(&input)

		choice, _ := strconv.Atoi(input)

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
