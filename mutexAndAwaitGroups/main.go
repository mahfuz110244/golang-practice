package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	fmt.Println("Race condition")
// 	wg := &sync.WaitGroup{}
// 	wg.Add(4)
// 	mut := &sync.RWMutex{}
// 	var score = []int{0}

// 	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
// 		fmt.Println("One 1")
// 		mut.Lock()
// 		score = append(score, 1)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)

// 	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
// 		fmt.Println("Two 2")
// 		mut.Lock()
// 		score = append(score, 2)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)

// 	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
// 		fmt.Println("Three 3")
// 		mut.Lock()
// 		score = append(score, 3)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)

// 	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
// 		fmt.Println("RLock test")
// 		mut.RLock()
// 		fmt.Println(score)
// 		mut.RUnlock()
// 		wg.Done()
// 	}(wg, mut)

// 	wg.Wait()
// 	fmt.Println(score)

// }

func main() {
	fmt.Println("Race condition")
	wg := &sync.WaitGroup{}
	wg.Add(3)
	mut := &sync.Mutex{}
	var score = []int{0}

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
