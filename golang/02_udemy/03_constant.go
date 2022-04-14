package main

import "fmt"

func main() {

	// Default values for types
	var name string
	var number int
	var FloatNumber float32
	var b bool

	fmt.Println("string default val ", name)
	fmt.Println("int default val ", number)
	fmt.Println("float default val ", FloatNumber)
	fmt.Println("boolean default val ", b)

	fmt.Println()

	// Constants
	const gameName = "game"
	fmt.Println("game name: ", gameName)

	// Constants cannot be reassigned
	// gameName = "new game"

}
