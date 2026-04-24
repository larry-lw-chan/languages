package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func type_main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Type Conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// String to Int
	n, err := strconv.Atoi("25")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	// Type Inference
	v := 45.5
	fmt.Printf("v is of type %T\n", v)
}
