package main

import (
	"fmt"
)

type Circle struct {
	radius float64
	area   float64
}

type Circle2 struct {
	radius float64
	area   float64
}

func (c *Circle) calcArea() float64 {
	// Method Intro
	// A method is a function with a receiver.

	// Syntax
	// func (<receiver>) <method name>(<params>) <return type> {
	// 	<body>
	// }

	// c is a pointer to the instance of circle struct
	c.area = 3.14 * c.radius * c.radius
	return c.area
}

func (c Circle2) calcAreaByValue() {
	// c is a NOT a pointer to the instance of circle struct
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	c := Circle{radius: 5}
	c.calcArea()
	fmt.Println(c.area)

	c2 := Circle2{radius: 5}
	c2.calcAreaByValue()
	fmt.Printf("%+v", c2)
}
