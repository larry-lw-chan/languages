package main

import "fmt"

// An empty interface may hold values of any type.
// (Every type implements at least zero methods.)

// Empty interfaces are used by code that handles values
// of unknown type. For example, fmt.Print takes any
// number of arguments of type interface{}.

func empty_interface_main() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
