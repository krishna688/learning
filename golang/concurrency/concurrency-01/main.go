package main

import (
	"fmt"
)

func sum(input []int, ch chan int) {
	var sum int

	for _, i := range input {
		sum += i
	}
	ch <- sum
}

func main() {

	input := []int{1, 4, 5, 6, 7}
	ch := make(chan int)

	go sum(input[:2], ch)
	go sum(input, ch)
	go sum1(nil, ch)

	x, y, z := <-ch, <-ch, <-ch

	fmt.Println(x, y, z)
}

func sum1(input []int, ch chan int) {
	var sum int

	for _, i := range input {
		sum += i
	}
	ch <- sum
}
