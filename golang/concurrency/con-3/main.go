package main

import (
	"fmt"
	"time"
)

func main() {

	tick := make(chan int)
	boom := time.After(6 * time.Second)

	go func() {
		for i := 5; i > 0; i-- {
			tick <- i
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case v := <-tick:
			fmt.Printf("tick %v", v)
		case <-boom:
			// time.Sleep(500 * time.Millisecond)
			fmt.Println("BOOM")
			return
		default:
			fmt.Println("     ...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
