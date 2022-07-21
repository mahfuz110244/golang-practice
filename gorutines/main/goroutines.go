// package main

// import (
// 	"fmt"
// 	"time"
// )

// func f(from string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(from, ":", i)
// 	}
// }

// func main() {
// 	// Direct Func call
// 	f("direct")

// 	// Golang go routines call
// 	go f("gorutines")

// 	// Go anonymmous function call
// 	go func(msg string) {
// 		fmt.Println(msg)
// 	}("going")

// 	time.Sleep(time.Second)

// 	fmt.Println("Done")
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	go countFruit("Apple")
	go countFruit("Orange")
	time.Sleep(time.Second * 10)
}

func countFruit(name string) {
	for i := 1; true; i++ {
		fmt.Println(name)
		time.Sleep(time.Second * 1)
	}
}
