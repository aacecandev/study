package main

import (
	"fmt"
	"time"
)

func main() {
	name := "foo"

	switch name {
	case "foo":
		fmt.Println("name is foo")
		fallthrough
	case "bar":
		fmt.Println("name is bar")
	default:
		fmt.Println("name is not foo")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
