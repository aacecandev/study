package main

import (
	"fmt"
)

func main() {
	var user string
	user = "Harry"
	fmt.Println(user)

	//incorrect
	/*
		var s string
		s = 123
	*/

	var s, t string = "foo", "bar"
	fmt.Println(s, t)

	x := "Hello"
	x = "World"
	fmt.Println(x)

	/*
		Error

		y := "Hello"
		y = 123
	*/
}
