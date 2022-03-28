package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i int = 9
	var f float64 = float64(i)

	fmt.Printf("%.2f\n", f)

	var x int = int(f)
	fmt.Printf("%d\n", x)

	var s string = strconv.Itoa(i)
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Println(reflect.TypeOf(s))

	x, err := strconv.Atoi(s)
	fmt.Printf("%v, %T \n", x, x)
	fmt.Printf("%v, %T", err, err)

	var errorString string = "200abc"
	y, err := strconv.Atoi(errorString)
	fmt.Printf("%v, %T \n", y, y)
	fmt.Printf("%v, %T", err, err)

}
