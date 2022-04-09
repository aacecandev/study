package main

import (
	"fmt"
)

func main() {
	// & operator
	// The address of a variable can be obtained by using the address-of operator.

	// * operator
	// The value at an address can be obtained by using the dereference operator.

	x := 77
	fmt.Printf("%T %v \n", &x, &x)
	fmt.Printf("%T %v \n", *(&x), *(&x))

}
