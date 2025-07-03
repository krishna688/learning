// Question 5: Go Routines with Channels
// Create a producer-consumer system:

// A producer goroutine sends integers (1 to 10) to a channel.
// A consumer goroutine reads from the channel and prints the values.

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}

}

func consumer(ch chan int, wg *sync.WaitGroup) {

	for i := range ch {
		fmt.Println(i)
	}
}
