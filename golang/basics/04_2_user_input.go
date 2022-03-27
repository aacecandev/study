package main

import (
	"fmt"
)

func main() {
	var a string
	var b int

	fmt.Print("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You entered: ", count, "values")
}
