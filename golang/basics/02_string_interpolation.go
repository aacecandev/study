package main

import "fmt"

func main() {
	var city string = "London"
	user := "Bob"
	fmt.Print("Welcome to ", city, ", ", user, "\n")

	fmt.Print(user, "\n")
	fmt.Println(user)

	fmt.Print("\n")

	fmt.Printf("Welcome to %s, %s\n", city, user)

	// %v - formats the value in a default format
	fmt.Printf("Welcome to %v, %v\n", city, user)

	// %d - formats decimal integers
	var grades int = 42
	fmt.Printf("%d\n", grades)

	// Verb - Description
	// %v - default format
	// %T - type of the value
	// %t - boolean
	// %b - binary
	// %c - character
	// %d - integers
	// %o - octal
	// %q - quoted character/string
	// %f - floating point
	// %.2f - floating point up to 2 decimals
	// %s - plain string
	// %c - character
	// %p - pointer
}
