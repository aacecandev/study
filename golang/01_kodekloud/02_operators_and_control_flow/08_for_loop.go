package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println(i)
	// }

	fmt.Println()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			break
		}
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println(i)
			continue
		}
	}
}
