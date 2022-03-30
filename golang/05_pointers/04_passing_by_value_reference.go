package main

import (
	"fmt"
)

func modifyValue(s string) {
	fmt.Println(&s)
	s = "b"
}

func modifyReference(d *string) {
	fmt.Println(d)
	*d = "d"
}

func main() {
	// passing by value
	// the parameter is copied into another location of your RAM
	// the original parameter is not changed, just the copy

	// all basic types (int, float, bool, string, array) are passed by value

	a := "a"
	fmt.Println(&a)
	fmt.Println(a)
	modifyValue(a)
	fmt.Println(a)

	// passing by reference
	// in call by reference/pointer, the address of the variable is passed
	// operations are performed on the content of that address

	// slices are passed by reference by default
	// maps are passed by reference by default

	c := "c"
	fmt.Println(&c)
	fmt.Println(c)
	modifyReference(&c)
	fmt.Println(c)
}
