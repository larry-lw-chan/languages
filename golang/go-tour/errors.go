package main

import (
	"fmt"
	"time"
)

// The error type is a built-in interface similar to fmt.Stringer:
// This shows how to implement the error interface using a custom type

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func error_main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
