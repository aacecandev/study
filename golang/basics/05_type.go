package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 42
	var b string = "a"
	var c bool = true
	var d float32 = 5.444

	fmt.Printf("The type of var a is: %T", a)
	fmt.Printf("\nThe type of var b is: %T", b)
	fmt.Printf("\nThe type of var c is: %T", c)
	fmt.Printf("\nThe type of var d is: %T", d)

	e := fmt.Sprintf("\nThe type of var a is: %T", a)
	fmt.Println(e)

	fmt.Printf("Type: %v \n", reflect.TypeOf(a))

	fmt.Printf("variable: %v is of type %v \n", b, reflect.TypeOf(b))
}
