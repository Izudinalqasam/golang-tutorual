package main

import "testing"

func BenchmarkFibonnaci(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
	fibonacci(2)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
