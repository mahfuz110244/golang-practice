// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go countFruit("Apple")
// 	go countFruit("Orange")
// 	time.Sleep(time.Second * 10)
// }

// func countFruit(name string) {
// 	for i := 1; true; i++ {
// 		fmt.Println(name)
// 		time.Sleep(time.Second * 1)
// 	}
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string)
	fmt.Println(ch)

	go countFruit("Apple", ch)
	ch <- "go man go"
	time.Sleep(time.Second * 5)
	fmt.Println(<-ch)

	go countFruit("orange", ch)
	ch <- "go man go orange"
	time.Sleep(time.Second * 5)
	fmt.Println(<-ch)
	fmt.Println(ch)
}

func countFruit(name string, ch chan string) {
	fmt.Println(ch)
	data := <-ch
	fmt.Println(name)
	fmt.Println(data)
	ch <- "stopped " + name
}
