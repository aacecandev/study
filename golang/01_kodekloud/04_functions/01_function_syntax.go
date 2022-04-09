package main

import (
	"fmt"
)

/*
	Naming convention:
		- <function_name> := <return_type> <function_name>(<parameters>)
		- <function_name> := <return_type> func(<parameters>)
		- must begin with a letter
		- can contain letters, numbers, and symbols
		- cannot contain spaces
		- must be unique
		- cannot be a reserved keyword
		- cannot be a predefined name
		- case-sensitive
*/

func main() {
	sumOfNumbers := addNumbers(1, 2)

	fmt.Println(sumOfNumbers)
}

// func <function_name>(<params>) <return type> { -> function signature
// function body
// }

func addNumbers(a int, b int) int {
	return a + b
}
