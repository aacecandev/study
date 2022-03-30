package main

import (
	"fmt"
)

func main() {
	x := func(a int, b int) int {
		return a * b
	}

	fmt.Printf("%T\n", x)
	fmt.Println(x(1, 2))

	y := func(a, b int) int {
		return a + b
	}(20, 30)

	fmt.Printf("%T\n", y)
	fmt.Println(y)
}
