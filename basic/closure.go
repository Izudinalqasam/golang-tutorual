package basic

import "fmt"

// Closures is a function which interact with data around it in same scope (function inside function can be called clouser in simple example)
// variable inside can't be accessed from outside, but outside can be accessed from inside function
func MainClosure() {
	fmt.Println(counterClosure())
}

func counterClosure() int {
	counter := 0

	increment := func() {
		counter++
	}

	increment()
	increment()
	return counter
}
