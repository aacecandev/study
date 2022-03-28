package main

import (
	"fmt"
)

func main() {
	var a, b string = "a", "b"
	fmt.Println(a == b)
	fmt.Println(a != b)

	var c, d int = 5, 6
	fmt.Println(c < d)
	fmt.Println(c > d)
	fmt.Println(c <= d)
	fmt.Println(c >= d)
}
