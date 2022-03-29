/*

fixed length
elements with same data type

every element has a memory address assigned to it

an array has a pointer that points to the first element

length is the number of elements in the array
length == capacity

*/

package main

import (
	"fmt"
)

func main() {
	var grades [5]int
	fmt.Println(grades)
	// [0 0 0 0 0]

	var grades2 [3]int = [3]int{1, 2, 3}
	fmt.Println(grades2)

	grades3 := [4]int{1, 2, 3, 4}
	fmt.Println(len(grades3))
	grades4 := [...]int{1, 2, 3}
	fmt.Println(len(grades4))

	fmt.Println(grades4[2])

	grades4[2] = 10
	fmt.Println(grades4[2])

	for i := 0; i < len(grades4); i++ {
		fmt.Println(grades4[i])
	}

	for index, element := range grades4 {
		fmt.Println(index, "=>", element)
	}

	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D[1][1])
}
