package main

import "fmt"

// Golang's Enum Approximation
type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func enum_main() {
	// Declare Enum
	myDirection := West

	switch myDirection {
	case North:
		fmt.Println("I'm going North")
	case South:
		fmt.Println("I'm going South")
	case East:
		fmt.Println("I'm going East")
	case West:
		fmt.Println("I'm going West")
	}
}
