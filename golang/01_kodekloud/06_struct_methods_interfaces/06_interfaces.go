package main

// An interface specifies a method set and is a powerful way to introduce modularity in Go
// An interface is like a blueprint for a method set
// They describe all the methods of a method set by providing the function signature for each method
// They specify the method, but not implement it

// type <interface_name> interface {
// method signatures
// }

type FixedDeposit interface {
	getRateOfInterest() float64
	calcReturn() float64
}

// implementing an interface

// a type, like a struct, implements an interface by implementing its methods
// Go interfaces are implemented implicitly
// Go hasn't any specific keyword to say that a type implements an interface

func main() {

}
