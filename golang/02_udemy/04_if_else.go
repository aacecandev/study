package main

import (
	"fmt"
)

func main() {
	name := "foo"

	if name == "foo" {
		fmt.Println("name is foo")
	} else if name == "bar" {
		fmt.Println("name is bar")
	} else {
		fmt.Println("name is not foo")
	}

	if true {
		fmt.Println("true")
	}

	if false {
		fmt.Println("false")
	}

	age := 18

	if age >= 18 && age <= 65 {
		if name == "bar" || name == "foo" {
			fmt.Println("age is 18 or older", age)
		}
	} else {
		fmt.Println("age is less than 18")
	}
}
