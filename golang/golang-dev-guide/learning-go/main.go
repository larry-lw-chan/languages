package main

import (
	"errors"
	"fmt"
)

func main() {
	num1 := 32.0
	num2 := 16.0

	result, err := divide(num1, num2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func divide(x,y float64) (float64, error) {
	var result float64
	
	if y == 0 {
		return result, errors.New("Cannot divide by zero")
	}
	result = x / y
	return result, nil
}