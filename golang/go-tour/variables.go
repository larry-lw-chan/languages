package main

import "fmt"

// A var statement can be at package or function level.
var c, python, java bool
var i, j int = 1, 2

func var_main() {
	var a int
	fmt.Println(a, c, python, java)
	fmt.Println(i + j)

	b := 3
	c := 5

	fmt.Println(b + c)
}
