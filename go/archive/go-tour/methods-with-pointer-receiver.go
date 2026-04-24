package main

import (
	"fmt"
	"math"
)

type Vertex6 struct {
	X, Y float64
}

func (v *Vertex6) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex6) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// There are two reasons to use a pointer receiver.
// The first is so that the method can modify the value
// that its receiver points to.

// The second is to avoid copying the value on each method
// call. This can be more efficient if the receiver is a
// large struct, for example.

func methods_with_pointer_receiver_main() {
	v := &Vertex6{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
