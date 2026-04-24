package main

import (
	"fmt"
	"math"
)

// Example 1 - Receiver Functions on Structs
type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Example 2 - Using Receiver Functions on regular Types
type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods_main() {
	// Example 1 - Using Receiver Functions on Structs
	v := Vertex4{3, 4}
	fmt.Println(v.Abs())
	fmt.Println()

	// Example of non receiver function working on struct
	fmt.Println(Abs2(v))
	fmt.Println()

	// Example 2 - Using Receiver Functions on regular Types
	f := myFloat(-10.0)
	fmt.Println(f.Abs())
}
