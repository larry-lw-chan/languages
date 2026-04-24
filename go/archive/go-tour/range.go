package main

import "fmt"

var pow3 = []int{1, 2, 4, 8, 16, 32, 64, 128}

func range_main() {
	// Example 1
	for i, v := range pow3 {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	println()
	println()

	// Example 2
	pow4 := [8]int{}

	for i := range pow4 {
		pow4[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow4 {
		fmt.Printf("%d\n", value)
	}
}
