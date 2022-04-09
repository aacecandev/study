package main

import (
	"fmt"
)

// factorial(5) = 5*4*3*2*1

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	n := 5
	fmt.Println(factorial(n))
}
