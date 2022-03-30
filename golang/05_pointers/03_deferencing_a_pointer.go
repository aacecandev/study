package main

import (
	"fmt"
)

func main() {
	s := "Hello"
	fmt.Printf("%T %v \n", s, s)

	ps := &s
	*ps = "World"

	fmt.Printf("%T %v \n", s, s)
}
