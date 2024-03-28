package main

import "fmt"

func slice_main() {
	// Example 1
	primes := [6]int{2, 3, 5, 7, 11, 13}
	for _, v := range primes {
		fmt.Println(v)
	}
	slice := primes[2:4]
	fmt.Println("This is\n\n", slice)

	// Example 2
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s, "\n\n")

	// Example 3
	a := []int{2, 3, 5, 7, 11, 13}

	a = a[1:4]
	fmt.Println(a)

	a = a[:2]
	fmt.Println(a)

	a = a[1:]
	fmt.Println(a)
}
