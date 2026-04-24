package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func higher_order_main() {
	// Example 1
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// Pass function into compute
	fmt.Println(compute(hypot))
	// Pass another function into compute
	fmt.Println(compute(math.Pow))
}
