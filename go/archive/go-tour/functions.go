package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// This shows function shortcuts
func subtract(x, y int) int {
	return x - y
}

// This is swap function
func swap(x, y string) (string, string) {
	return y, x
}

// This shows named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func func_main() {
	fmt.Println(add(1, 2))
	fmt.Println(subtract(5, 2))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
}
