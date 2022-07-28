// Do not communicate by sharing memory; instead share memory by communicating
Let me go simple and to the point here.

Don't communicate by sharing memory.

It is like when your are communicating using threads for example you have to use variables or mutexes to lock memory for not allowing someone to read and write on it until the communication is complete.

Share memory by communicating

In go routines values move on channels rather than blocking the memory, sender notifies receiver to receive from that channel and hence it share memory by communicating with receiver to get from a channel.

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println((s))
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
