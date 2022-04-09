package main

import (
	"fmt"
)

func main() {
	// continuous segment of an underlying array

	// variable typed (elements can be added or removed)

	// more flexible

	/*

		[0,1,2,3,4,5]
		  [0,1,2,3]

		elements:
		  pointer to the first element of the underlying array 1 -> 0
			length: number of elements the slice contains -> 4
			capacity: number of elements in the underlying array counting
			  from the first element on the slice -> [1,2,3,4,5] of the underlying array
	*/

	// <slice_name> := []<data_type>{<values}
	// slice := make([]<data_type>, length, capacity)

	slice := []int{10, 20, 30}
	fmt.Println(slice)

	// array[start_included: end_excluded]

	arr := [5]int{1, 2, 3, 4, 5}

	slice1 := arr[1:3]
	fmt.Println(slice1)

	subSlice := slice1[1:2]
	fmt.Println(subSlice)

	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// if an element of the slice is changed, the underlying array element is changed too

	arr1 := [5]int{1, 2, 3, 4}
	sliceArr1 := arr1[1:3]
	sliceArr1[0] = 100

	fmt.Println("arr1")
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))

	fmt.Println("sliceArr1")
	fmt.Println(sliceArr1)
	fmt.Println(len(sliceArr1))
	fmt.Println(cap(sliceArr1))

	// appending to a slice

	/*
		func append(slice []Type, values of that ...Type) []T result is a slice of []Type containing those values

		func append(s []T, v ...T) []T

		If we allocate more values than expected in the underlying array, the underlying array is extended

	*/

	slice2 = append(sliceArr1, 88, 99, 101)

	fmt.Println("slice2")
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2)) // The slice doubles its capacity

	// appending a slice to another slice

	// slice = append(slice, slice2...)

	arr2 := [5]int{10, 20, 30, 40, 50}
	sliceArr2 := arr2[:2]

	arr_2 := [5]int{5, 15, 25, 35, 45}
	slice_2 := arr_2[:2]

	new_slice := append(sliceArr2, slice_2...)

	fmt.Println("new_slice")
	fmt.Println(new_slice)

	// deleting from a slice

	arr3 := [5]int{10, 20, 30, 40, 50}
	i := 2

	fmt.Println("arr3")
	fmt.Println(arr3)

	slice3_1 := arr3[:i]   // 10, 20
	slice3_2 := arr3[i+1:] // 40, 50

	new_slice3 := append(slice3_1, slice3_2...)

	fmt.Println("new_slice3")
	fmt.Println(new_slice3)

	// copying from a slice

	// func copy(dst, src []Type) int
	// returns the number of elements copied, the minimum of the length of the destination slice or the source slice

	// The slices must be initalized with the same data type

	// num := copy(new_slice3, slice3_1)

	src_slice := []int{10, 20, 30, 40, 50}
	dst_slice := make([]int, 3)

	num := copy(dst_slice, src_slice)

	fmt.Println("dest_slice")
	fmt.Println(dst_slice)
	fmt.Println("num of elements copied")
	fmt.Println(num)

	// looping through a slice

	for index, value := range dst_slice {
		fmt.Println(index, value)
	}

}
