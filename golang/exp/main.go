package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int) int {
	time.Sleep(1 * time.Second)
	return i
}

func main() {

	dataCh := make(chan int)

	go func() {

		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dataCh <- doSomething(i)
			}()
		}

		wg.Wait()
		close(dataCh)
	}()

	for n := range dataCh {
		fmt.Println(n)
	}

}
