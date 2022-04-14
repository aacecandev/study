package main

import "fmt"

func main() {
	// Declare variables  without assigning values
	var name string
	name = "John"

	fmt.Println("var name: ", name)

	// Reassign variables

	var name2 string
	name2 = "John"
	fmt.Println("var name2: ", name2)

	name2 = "Jane"
	fmt.Println("var name2: ", name2)

	// Declare variables without assigning types
	var name3 = "John"
	fmt.Println("var name3: ", name3)

	name3 = "Jane"
	fmt.Println("var name3: ", name3)

	// Redeclaration of variables is not allowed
	// var VarName = "data"
	// fmt.Println(VarName)

	// var VarName = "new data"
	// fmt.Println(VarName)
}
