package main

import "fmt"

func main() {
	fmt.Println("working")
	fmt.Println(Fibonacci(5))
}

func Fibonacci(n int) []int {
	// Your code here
	fib := []int{1, 2}

	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib
}
