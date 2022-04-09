package main

import (
	"fmt"
)

type s1 struct {
	x int
}

type s2 struct {
	x int
}

func main() {
	// Error since they are not the same type

	// c := s1{x: 5}
	// c1 := s2{x: 5}

	// if c == c1 {
	// 	fmt.Println("c == c1")
	// } else {
	// 	fmt.Println("c != c1")
	// }

	c2 := s1{x: 5}
	c3 := s1{x: 6}

	if c2 == c3 {
		fmt.Println("c2 == c3")
	} else {
		fmt.Println("c2 != c3")
	}
}
