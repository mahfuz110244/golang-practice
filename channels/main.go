package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Golang Channels")
	myCh := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, ischannelOpen := <-ch
		fmt.Println(val)
		fmt.Println(ischannelOpen)
		wg.Done()

	}(myCh, wg)

	// Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// close(ch)
		ch <- 0
		close(ch)
		wg.Done()

	}(myCh, wg)
	wg.Wait()
}
