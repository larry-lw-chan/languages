package main

import "fmt"

func for_main() {
	// Loop Example 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Loop Example 2
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Loop Example 3 - Infinite Loop
	for {
		fmt.Println("Infinite Loop")
	}
}
