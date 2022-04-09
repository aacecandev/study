package main

import (
	"fmt"
)

// Method sets
// A method set is a set of methods that belong to a data type.
// Useful for encapsulating functionality

type Student struct {
	name   string
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func main() {
	s := Student{name: "John", grades: []int{1, 2, 3, 4, 5}}
	s.displayName()
	fmt.Printf("%.2f", s.calculatePercentage())
}
