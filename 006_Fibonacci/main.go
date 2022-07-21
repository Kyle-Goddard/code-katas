package main

import "fmt"

func main() {
	idx := 1

	fmt.Printf("Fibonacci number %v: %v\n", idx, fib(idx))
}

func fib(n int) int {
	a, b, i, f := 0, 1, 1, 0

	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}
	for i < n {
		f = a + b

		i++
		a = b
		b = f
	}

	return f
}
