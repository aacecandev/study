package main

import (
	"fmt"
)

// type <struct_name> struct { list of fields }

type Circle struct {
	x float64
	y float64
	r float64
}

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

func main() {

	// var <variable_name> <struct_name>

	var s Student
	fmt.Printf("%+v\n", s)

	s.name = "John"
	s.rollNo = 1
	s.marks = []int{10, 20, 30}
	s.grades = map[string]int{"A": 10, "B": 20}

	// this returns the memory address
	st := new(Student)
	fmt.Printf("%+v\n", st)

	sc := Student{
		name:   "John",
		rollNo: 1,
		marks:  []int{10, 20, 30},
		grades: map[string]int{"A": 10, "B": 20},
	}

	fmt.Printf("%+v\n", sc)

	sd := Student{"Joe", 1, []int{10, 20, 30}, map[string]int{"A": 10, "B": 20}}
	fmt.Printf("%+v\n", sd)
}
