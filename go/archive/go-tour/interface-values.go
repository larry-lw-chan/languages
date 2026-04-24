package main

import (
	"fmt"
	"math"
)

type I2 interface {
	M()
}

// Type T2 implements the interface I2
type T2 struct {
	S string
}

func (t *T2) M() {
	fmt.Println(t.S)
}

// Type F implements the interface I2
type F float64

func (f F) M() {
	fmt.Println(f)
}

// Type G implements the interface I2 with nil
type G struct {
	S string
}

func (g *G) M() {
	if g == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(g.S)
}

func interface_3_main() {
	var i I2
	// Call T2
	i = &T2{"Hello"}
	describe2(i)
	i.M()
	fmt.Println()

	// Call I
	i = F(math.Pi)
	describe2(i)
	i.M()
	fmt.Println()

	// Call F with nil
	var g *G
	i = g
	describe2(i)
	i.M()
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
