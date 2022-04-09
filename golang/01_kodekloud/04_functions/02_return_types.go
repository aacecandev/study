package main

import (
	"fmt"
	"strings"
)

func operation(a int, b int) (int, int) {
	sum := a + b
	dif := a - b
	return sum, dif
}

func operation2(a int, b int) (sum int, dif int) {
	sum = a + b
	dif = a - b
	return
}

// variadic functions
// func <function_name>(<p1> <p1 type>, <p2> <p2 type> ...type) <return type> {}
func sumNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// variadic function 2
func printDetails(student string, subjects ...string) {
	fmt.Println("Hey", student, ", here are your subjects: ")
	for _, subject := range subjects {
		fmt.Printf("%s ", subject)
	}
	return
}

func printStrings(names ...string) (names_c []string) {
	names_c = []string{}
	for _, value := range names {
		names_c = append(names_c, strings.ToUpper(value))
	}
	return
}

func main() {
	sum, dif := operation(1, 2)
	sum2, dif2 := operation2(1, 2)
	fmt.Println(sum, dif)
	fmt.Println(sum2, dif2)

	fmt.Println(sumNumbers(1, 2, 3, 4, 5))

	printDetails("John", "Maths", "Physics", "Chemistry")

	result := printStrings("Joe", "Monica", "Gunther")
	fmt.Println(result)
}
