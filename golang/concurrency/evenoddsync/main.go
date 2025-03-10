package main

import (
	"log"
	"sync"
)

func printEvent(input int, evChan, odChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= input; {
		<-evChan
		log.Println(i)
		i += 2
		odChan <- 1
	}
}

func main() {
	input := 10
	wg := sync.WaitGroup{}
	evCh := make(chan int, 1)
	odCh := make(chan int, 1)

	wg.Add(2)
	go printEvent(input, evCh, odCh, &wg)
	go printOdd(input, evCh, odCh, &wg)
	odCh <- 1

	wg.Wait()
}

func printOdd(input int, evCh, odChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= input; {
		<-odChan
		log.Println(i)
		i += 2
		evCh <- 1
	}
}
