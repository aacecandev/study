package main

import (
	"fmt"
)

func PowersOfTwo(n int) []uint64 {
	list := []uint64{1}

	for i := 1; i <= n; i++ {
		list = append(list, uint64(2)*list[i-1])
	}

	return list

}

func main() {
	fmt.Println(PowersOfTwo(10))
}
