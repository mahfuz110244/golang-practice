package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Direct Func call
	f("direct")

	// Golang go routines call
	go f("gorutines")

	// Go anonymmous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)

	fmt.Println("Done")
}
