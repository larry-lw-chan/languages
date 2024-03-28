package main

import (
	"fmt"
	"math"
)

// Example 1
func pow(x, n, lim float64) float64 {
	// Note V is in the scope of the if statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Example 2
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func if_main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
