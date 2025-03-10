package main

import (
	"log"
	"sync"
	"sync/atomic"
)

var (
	count int64
)

func getPrimenumbers(ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	for val := range ch {
		if isPrimeNumber(val) {
			atomic.AddInt64(&count, 1)
			log.Println(val)
		}

	}

}

func isPrimeNumber(num int) bool {

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	start := 10
	end := 50
	n := 4

	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go getPrimenumbers(ch, &wg, i)
	}
	go func() {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Wait()

	log.Printf("total count %d", count)
}
