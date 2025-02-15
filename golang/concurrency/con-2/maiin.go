package main

import "fmt"

func fibnocci(n int, ch chan int) {

	x, y := 0, 1

	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {

	ch := make(chan int)
	go fibnocci(10, ch)

	for val := range ch {
		fmt.Println(val)
	}
}
