package main

import (
	"fmt"
)

func main() {
	var name string
	var isSmart bool = false

	fmt.Print("Enter your name & you're smart: ")
	fmt.Scanf("%s %t", &name, &isSmart)
	fmt.Println("Hello, ", name, "! You're", isSmart)
}
