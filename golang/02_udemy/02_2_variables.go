package main

import "fmt"

func main() {

	// Declare variables using the shorthand declaration operator
	name := "John"
	fmt.Println("var name: ", name)

	// Error, = is not an assignment operator
	// name2 = "John"
	// fmt.Println("var name: ", name2)

	// Changing a varible type throws an error
	// name := 1

	// Types
	game := "XO"
	players := 2
	gameStarted := true

	fmt.Println("var game: ", game, "players: ", players, "gameStarted: ", gameStarted)

	fmt.Printf("var game: ", game, "players: ", players, "gameStarted: ", gameStarted)
	fmt.Println()
	fmt.Printf("var game: %s players: %d gameStarted: %t\n", game, players, gameStarted)
}
