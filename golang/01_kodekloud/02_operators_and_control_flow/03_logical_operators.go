package main

import (
	"fmt"
)

func main() {
	// AND &&
	// true if both operands are true
	// false if either operand is false

	var x int = 10

	fmt.Println(x > 5 && x < 15)
	fmt.Println(x > 15 && x < 5)

	// OR ||
	// true if either operand is true
	// false if both operands are false

	fmt.Println(x > 5 || x < 15)
	fmt.Println(x > 15 || x < 5)

	// NOT !
	// true if operand is false
	// false if operand is true

	fmt.Println(!(x > 5 && x < 15))
	fmt.Println(!(x > 15 && x < 5))
}
