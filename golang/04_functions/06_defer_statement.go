package main

import (
	"fmt"
)

func printName(name string) {
	fmt.Println(name)
}

func printRollNo(role int) {
	fmt.Println(role)
}

func printAddress(address string) {
	fmt.Println(address)
}

func main() {
	// A defer statement delays the execution of a function until the surrounding function returns.

	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	printName("Joe")
	defer printRollNo(23)
	printAddress("street")
}
