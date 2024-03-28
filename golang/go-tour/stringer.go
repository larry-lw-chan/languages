package main

import "fmt"

// A Stringer is a type that can describe itself as a string.
// Any type that implements this can use the fmt package .
// FMT (and many others) look for this interface to print values

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringer_main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a)
	fmt.Println(z)
}
