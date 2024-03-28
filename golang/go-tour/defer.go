package main

import "fmt"

func defer_main() {
	// Defer
	// defer fmt.Println("world")
	// fmt.Println("hello")

	// Stacking Defers
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
