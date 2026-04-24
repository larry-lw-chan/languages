package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf(": cannot Sqrt negative number: %v", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	if x >= 0 {
		z := 1.0
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	} else {
		return x, ErrNegativeSqrt(x)
	}
}

func error_exercise_main() {
	fmt.Println(Sqrt2(2.0))
	fmt.Println(Sqrt2(-2.0))
}
