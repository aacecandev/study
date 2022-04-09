package main

import (
	"fmt"
)

func main() {
	// var <pointer_name> *<type>

	var ptr_i *int
	var ptr_s *string

	fmt.Println(ptr_i)
	fmt.Println(ptr_s)

	fmt.Println()

	// var <pointer_name> *<type> = &<variable_name>

	i2 := 10
	var ptr_i2 *int = &i2

	fmt.Println("i2")
	fmt.Println(ptr_i2)
	fmt.Println(*ptr_i2)

	fmt.Println()

	// var <pointer_name> = &<variable_name>

	s2 := "s2"
	var ptr_s2 = &s2

	fmt.Println("s2")
	fmt.Println(ptr_s2)
	fmt.Println(*ptr_s2)

	fmt.Println()

	// <pointer_name> := &<variable_name>

	s3 := "s3"
	ptr_s3 := &s3

	fmt.Println("s3")
	fmt.Println(ptr_s3)
	fmt.Println(*ptr_s3)

}
