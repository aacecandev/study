package main

import (
	"fmt"
)

func main() {
	// bitwise AND (&)
	// takes two integers and performs AND on every bit of the two numbers
	// returns 1 when both bits are 1, else 0

	// 12 = 0000 1100
	// 25 = 0001 0101

	var a, b int = 12, 25

	z := a & b
	// 0000 1000
	fmt.Println(z)

	// bitwise OR (|)
	// takes two integers and performs OR on every bit of the two numbers
	// returns 1 when either bit is 1, else 0

	z = a | b
	// 0001 1101
	fmt.Println(z)

	// bitwise XOR (^)
	// takes two integers and performs XOR on every bit of the two numbers
	// returns 1 when only one of the bits is 1, else 0

	z = a ^ b
	// 0001 0101
	fmt.Println(z)

	// left shift (<<)
	// shifts bits to the left by the number of positions specified
	// returns the value of the left-shifted bits

	z = a << 1
	//  0000 1100
	// 0000 1100
	// 0000 11000
	fmt.Println(z)

	// right shift (>>)
	// shifts bits to the right by the number of positions specified
	// returns the value of the right-shifted bits

	z = a >> 1
	// 0000 1100
	//  0000 1100
	// 00000 1100
	fmt.Println(z)

	// bitwise NOT (~)
	// takes an integer and performs NOT on every bit of the numbers
	// returns 1 when the bit is 0, else 0

}
