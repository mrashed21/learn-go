package main

import "fmt"

func function() {
	// Named function
	testFunction := func() {
		fmt.Println("This is a test function.")
	}

	testFunction()

	// Anonymous function immediately invoked(IIFE)
	func() {
		fmt.Println("This is an anonymous function.")
	}()
}
