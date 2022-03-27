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

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var ch = make(chan string)
// 	fmt.Println(ch)

// 	go countFruit("Apple", ch)
// 	ch <- "go man go"
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(<-ch)

// 	go countFruit("orange", ch)
// 	ch <- "go man go orange"
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(<-ch)
// 	fmt.Println(ch)
// }

// func countFruit(name string, ch chan string) {
// 	fmt.Println(ch)
// 	data := <-ch
// 	fmt.Println(name)
// 	fmt.Println(data)
// 	ch <- "stopped " + name
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var ch = make(chan string)
// 	go countFruit(ch)
// 	time.Sleep(time.Second * 10)
// 	ch <- "stopped "

// 	fmt.Println(<-ch)
// }

// func countFruit(ch chan string) {
// 	fmt.Println(ch)
// 	data := <-ch
// 	fmt.Println(data)
// 	time.Sleep(time.Second * 1)
// 	fmt.Println(time.Now())

// 	ch <- "stopped me if you can"
// }

package main

import (
	"fmt"
	"time"
)

func goRoutine(ch chan string) {
	for {
		fmt.Println("working, now wait for 1 second in goroutine")
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now())
		data := <-ch
		fmt.Println(fmt.Sprintf("signal from main goroutine: %s", data))
		if data != "stop" {
			fmt.Println("working in goroutine")
		} else {
			fmt.Println("stopped in goroutine")
			return
		}
	}
}

func main() {
	fmt.Println(fmt.Sprintf("main goroutine starting: %s", time.Now()))
	ch := make(chan string)
	go goRoutine(ch)
	fmt.Println("wait for 20 second in main goroutine")
	time.Sleep(20 * time.Second)
	ch <- "stop"
	time.Sleep(1 * time.Second)
	fmt.Println(fmt.Sprintf("main goroutine finished: %s", time.Now()))
}
