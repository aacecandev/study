// Create a slice of strings.
// Contains a length and capacity of 5 elements.
slice := make([]string, 5)

// Create a slice of integers.
// Contains a length of 3 and has a capacity of 5 elements.
slice := make([]int, 3, 5)

// Create a slice of integers.
// Make the length larger than the capacity.
slice := make([]int, 5, 3)

Compiler Error:
len larger than cap in make([]int)

// Create a slice of strings.
// Contains a length and capacity of 5 elements.
slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

// Create a slice of integers.
// Contains a length and capacity of 3 elements.
slice := []int{10, 20, 30}

Provide the minimal code, using a slice literal, that would yield a slice of 2 pointers to integers, pointing respectively to values 2 and 5

i1 := 2; i2 := 5; slc := []*int{&i1, &i2}

// Create a slice of strings.
// Initialize the 100th element with an empty string.
slice := []string{99: ""}

Remember, if you specify a value inside the [ ] operator, you’re creating an array. If you don’t specify a value, you’re creating a slice

// Create an array of three integers.
array := [3]int{10, 20, 30}

// Create a slice of integers with a length and capacity of three.
slice := []int{10, 20, 30}

// Create a nil slice of integers.
var slice []int

A nil slice is the most common way you create slices in Go. They can be used with many of the standard library and built-in functions that work with slices. They’re useful when you want to represent a slice that doesn’t exist, such as when an exception occurs in a function that returns a slice 

// Use make to create an empty slice of integers.
slice := make([]int, 0)

// Use a slice literal to create an empty slice of integers.
slice := []int{}

An empty slice contains a zero-element underlying array that allocates no storage. Empty slices are useful when you want to represent an empty collection, such as when a database query returns zero results.

// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}

// Change the value of index 1.
slice[1] = 25

// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}


// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]

newSlice can’t access the elements of the underlying array that are prior to its pointer. As far as newSlice is concerned, those elements don’t even exist. 

Calculating the length and capacity for any new slice is performed using the following formula.
For slice[i:j] with an underlying array of capacity k

Length:   j - i
Capacity: k - i

For slice[1:3] with an underlying array of capacity 5

Length:   3 - 1 = 2
Capacity: 5 - 1 = 4

You need to remember that you now have two slices sharing the same underlying array. Changes made to the shared section of the underlying array by one slice can be seen by the other slice

// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}

// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]

// Change index 1 of newSlice.
// Change index 2 of the original slice.
newSlice[1] = 35

A slice can only access indexes up to its length. Trying to access an element outside of its length will cause a runtime exception. The elements associated with a slice’s capacity are only available for growth. They must be incorporated into the slice’s length before they can be used.

// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}

// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]

// Change index 3 of newSlice.
// This element does not exist for newSlice.
newSlice[3] = 45

Runtime Exception:
panic: runtime error: index out of range

To use append, you need a source slice and a value that is to be appended. When your append call returns, it provides you a new slice with the changes. The append function will always increase the length of the new slice. The capacity, on the other hand, may or may not be affected, depending on the available capacity of the source slice

// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}

// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]

// Allocate a new element from capacity.
// Assign the value of 60 to the new element.
newSlice = append(newSlice, 60)

When there’s no available capacity in the underlying array for a slice, the append function will create a new underlying array, copy the existing values that are being referenced, and assign the new value. 

// Create a slice of integers.
// Contains a length and capacity of 4 elements.
slice := []int{10, 20, 30, 40}

// Append a new value to the slice.
// Assign the value of 50 to the new element.
newSlice := append(slice, 50)

Capacity is always doubled when the existing capacity of the slice is under 1,000 elements. Once the number of elements goes over 1,000, the capacity is grown by a factor of 1.25, or 25%. This growth algorithm may change in the language over time. 

This third index gives you control over the capacity of the new slice. The purpose is not to increase capacity, but to restrict the capacity

// Create a slice of strings.
// Contains a length and capacity of 5 elements.
source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

// Slice the third element and restrict the capacity.
// Contains a length of 1 element and capacity of 2 elements.
slice := source[2:3:4]

For slice[i:j:k]  or  [2:3:4]

Length:   j - i  or  3 - 2 = 1
Capacity: k - i  or  4 - 2 = 2

the first value represents the starting index position of the element the new slice will start with—in this case, 2. The second value represents the starting index position (2) plus the number of elements you want to include (1); 2 plus 1 is 3, so the second value is 3. For setting capacity, you take the starting index position of 2, plus the number of elements you want to include in the capacity (2), and you get the value of 4. 

// This slicing operation attempts to set the capacity to 4.
// This is greater than what is available.
slice := source[2:3:6]

Runtime Error:
panic: runtime error: slice bounds out of range

By having the option to set the capacity of a new slice to be the same as the length, you can force the first append operation to detach the new slice from the underlying array. Detaching the new slice from its original source array makes it safe to change. 

// Create a slice of strings.
// Contains a length and capacity of 5 elements.
source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

// Slice the third element and restrict the capacity.
// Contains a length and capacity of 1 element.
slice := source[2:3:3]

// Append a new string to the slice.
slice = append(slice, "Kiwi")

you can pass multiple values to be appended in a single slice call. If you use the ... operator, you can append all the elements of one slice into another. 

// Create two slices each initialized with two integers.
s1 := []int{1, 2}
s2 := []int{3, 4}

// Append the two slices together and display the results.
fmt.Printf("%v\n", append(s1, s2...))

Output:
[1 2 3 4]

Iterating over slices

// Create a slice of integers.
// Contains a length and capacity of 4 elements.
slice := []int{10, 20, 30, 40}

// Iterate over each element and display each value.
for index, value := range slice {
  fmt.Printf("Index: %d  Value: %d\n", index, value)
}


Output:
Index: 0  Value: 10
Index: 1  Value: 20
Index: 2  Value: 30
Index: 3  Value: 40

The keyword range, when iterating over a slice, will return two values. The first value is the index position and the second value is a copy of the value in that index position

It’s important to know that range is making a copy of the value, not returning a reference. If you use the address of the value variable as a pointer to each element, you’ll be making a mistake

// Create a slice of integers.
// Contains a length and capacity of 4 elements.
slice := []int{10, 20, 30, 40}

// Iterate over each element and display the value and addresses.
for index, value := range slice {
   fmt.Printf("Value: %d  Value-Addr: %X  ElemAddr: %X\n",
       value, &value, &slice[index])
}

Output:
Value: 10  Value-Addr: 10500168  ElemAddr: 1052E100
Value: 20  Value-Addr: 10500168  ElemAddr: 1052E104
Value: 30  Value-Addr: 10500168  ElemAddr: 1052E108
Value: 40  Value-Addr: 10500168  ElemAddr: 1052E10C

// Create a slice of integers.
// Contains a length and capacity of 4 elements.
slice := []int{10, 20, 30, 40}

// Iterate over each element and display each value.
for _, value := range slice {
    fmt.Printf("Value: %d\n", value)
}

Output:
Value: 10
Value: 20
Value: 30
Value: 40

// Create a slice of integers.
// Contains a length and capacity of 4 elements.
slice := []int{10, 20, 30, 40}

// Iterate over each element starting at element 3.
for index := 2; index < len(slice); index++ {
    fmt.Printf("Index: %d  Value: %d\n", index, slice[index])
}

Output:
Index: 2  Value: 30
Index: 3  Value: 40

There are two special built-in functions called len and cap that work with arrays, slices, and channels. For slices, the len function returns the length of the slice, and the cap function returns the capacity

Multidimensional slices

// Create a slice of a slice of integers.
slice := [][]int{{10}, {100, 200}}

The outer slice contains two elements, each of which are slices. The slice in the first element is initialized with the single integer 10 and the slice in the second element contains two integers, 100 and 200. 

// Create a slice of a slice of integers.
slice := [][]int{{10}, {100, 200}}

// Append the value of 20 to the first slice of integers.
slice[0] = append(slice[0], 20)

When the operation is complete, an entire new slice of integers and a new underlying array is allocated and then copied back into index 0 of the outer slice

Passing slices between functions 

// Allocate a slice of 1 million integers.
slice := make([]int, 1e6)

// Pass the slice to the function foo.
slice = foo(slice)

// Function foo accepts a slice of integers and returns the slice back.
func foo(slice []int) []int {
    ...
    return slice
}

On a 64-bit architecture, a slice requires 24 bytes of memory. The pointer field requires 8 bytes, and the length and capacity fields require 8 bytes respectively. Since the data associated with a slice is contained in the underlying array, there are no problems passing a copy of a slice to any function. Only the slice is being copied, not the underlying array

	// Compute and print Pascal's Triangle up to the n-th degree
	// Row 'i' has i+1 elements
	// for k=0 and k=i : pascal[i][0] = 1
	// for 0 < k < i : pascal[i][k] = pascal[i-1][k-1] + pascal[i-1][k]
	const n = 15
	pascal := make([][]int,n)

	for i,r := range pascal {
		r = make([]int,i+1)
    r[0] = 1; r[i] = 1
		for j := 1; j < i; j++ {
			r[j] = pascal[i-1][j-1] + pascal[i-1][j]
		}
		pascal[i] = r
		fmt.Println(i,": ",pascal[i])
	}