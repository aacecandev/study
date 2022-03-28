package main

import (
	"fmt"
)

func main() {
	// assign (=)
	var a int = 1
	var b int

	b = a
	fmt.Println(b)

	// add and assign (+=)
	a += 1 // a = a + 1
	fmt.Println(a)

	// subtract and assign (-=)
	a -= 1 // a = a - 1
	fmt.Println(a)

	// multiply and assign (*=)
	a *= 2 // a = a * 2
	fmt.Println(a)

	// divide and assign (/=)
	a /= 2 // a = a / 2
	fmt.Println(a)

	// modulo and assign (%=)
	a %= 2 // a = a % 2
	fmt.Println(a)

}
