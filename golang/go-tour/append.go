package main

import "fmt"

func append_main() {
	var s []int
	printSlice4(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice4(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice4(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice4(s)
}

func printSlice4(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
