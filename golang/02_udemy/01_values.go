package main

import (
	"fmt"
)

func main() {
	fmt.Println("strings")
	fmt.Println("123 + 123")
	fmt.Println("hello", "1", "name")

	fmt.Println()

	fmt.Println("numbers")
	fmt.Println(1, 2)
	fmt.Println(4 / 2)

	fmt.Println()

	fmt.Println("strings and numbers")
	fmt.Println("string ", 2)
	fmt.Println(2, "string")

	fmt.Println()

	fmt.Println("float numbers")
	fmt.Println(1.2, 1.2)
	fmt.Println("7.0/2.0=", 7.0/2.0)
	fmt.Println("7.0/2=", 7.0/2)
	fmt.Println("7/2=", 7/2)

	fmt.Println()

	fmt.Println("booleans")
	fmt.Println(true, false)
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true && !false)
	fmt.Println(!true || !false)

}
