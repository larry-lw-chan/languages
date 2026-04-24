package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

// Example 1 - Using Pointer Receiver Functions on Structs

// Methods with pointer receivers can modify
// the value to which the receiver points
// (as Scale does here). Since methods often
// need to modify their receiver, pointer receivers are
// more common than value receivers.
func (v *Vertex5) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Example 2 - Pointer & Functions
func Scale2(v *Vertex5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Example 4 - Pointer & Indirections
func AbsFunc(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func method_pointer_main() {
	// Example 1 - Using Pointer Receiver Functions on Structs
	v := Vertex5{3, 4}
	fmt.Println(v.X, v.Y)
	v.Scale(10)
	fmt.Println(v.X, v.Y)
	fmt.Println()

	// Example 2 - Pointer & Functions
	v2 := Vertex5{3, 4}
	fmt.Println(v2.X, v2.Y)
	Scale2(&v2, 10)
	fmt.Println(v2.X, v2.Y)
	fmt.Println()

	// Example 3 - Pointer Receivers can receive
	// both values and pointers
	v3 := Vertex5{6, 8}
	fmt.Println(v3.X, v3.Y)

	// First throw value into receiver
	// Golang will automatically convert it to a pointer
	v3.Scale(10)
	fmt.Println(v3.X, v3.Y)

	// Now throw pointer into receiver
	v3Ptn := &v3
	v3Ptn.Scale(10)
	fmt.Println(v3.X, v3.Y)
	fmt.Println()

	// Example 4 - Methods and pointer indirection (2)
	v4 := Vertex5{3, 4}
	fmt.Println(v4.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex5{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
