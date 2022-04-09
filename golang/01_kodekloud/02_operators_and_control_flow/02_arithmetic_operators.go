package main

import (
	"fmt"
)

func main() {
	var a, b string = "a", "b"
	fmt.Println(a + b)
	// fmt.Println(a - b) Error invalid operation

	var c, d float64 = 5.02, 6.98
	fmt.Printf("%.02f\n", c-d)

	fmt.Printf("%.02f\n", c*d)

	fmt.Printf("%.02f\n", c/d)

	var e, f int = 24, 7
	fmt.Println(e % f)

	e++
	fmt.Println(e)
	e--
	fmt.Println(e)

}
