package main

import (
	"fmt"
)

func main() {
	var i int = 100

	switch i {
	case 10:
		fmt.Println("i is 10")
	case 20:
		fmt.Println("i is 20")
	case 100:
		fmt.Println("i is 100")
	default:
		fmt.Println("i is not 10, 20 or 100")
	}

	fmt.Println()

	switch {
	case 10 < 1:
		fmt.Println("i is 10")
	case 100 == i:
		fmt.Println("i is 100")
		fallthrough // fallthrough is optional
	case 10+10 == 21:
		fmt.Println("i is 20")
		fallthrough // fallthrough is optional
	default:
		fmt.Println("i is not 10, 20 or 100")
	}
}
