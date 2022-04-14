package main

import (
	"fmt"
)

func main() {
	var names [3]string
	fmt.Println("names: ", names)

	var nums [3]int
	fmt.Println("nums: ", nums)

	names[0] = "foo"
	nums[0] = 1

}
