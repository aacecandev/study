package main

import (
	"fmt"
)

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return s.side * 4
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {
	return (r.length + r.width) * 2
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rectangle{length: 10, width: 5}
	s := square{side: 5}

	printData(r)
	printData(s)
}
