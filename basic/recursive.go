package main

import "fmt"

func MainRecursive() {
	fmt.Println(factorialLoop(5))
}

func factorialLoop(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorialLoop(number-1)
	}
}
