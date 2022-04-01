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

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func goRoutine(ch chan string) {
// 	for {
// 		fmt.Println("working, now wait for 1 second in goroutine")
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(time.Now())
// 		data := <-ch
// 		fmt.Println(fmt.Sprintf("signal from main goroutine: %s", data))
// 		if data != "stop" {
// 			fmt.Println("working in goroutine")
// 		} else {
// 			fmt.Println("stopped in goroutine")
// 			return
// 		}
// 	}
// }

// func main() {
// 	fmt.Println(fmt.Sprintf("main goroutine starting: %s", time.Now()))
// 	ch := make(chan string)
// 	go goRoutine(ch)
// 	fmt.Println("wait for 20 second in main goroutine")
// 	time.Sleep(20 * time.Second)
// 	ch <- "stop"
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(fmt.Sprintf("main goroutine finished: %s", time.Now()))
// }

// Create a goroutine that will loop and sleep for 1 second. When it wakes up it will print out the time.
// The main routine will wait for 20 seconds and then send a termination signal to the goroutine.
package main

import (
	"fmt"
	"time"
)

func goRoutine(quit chan string) {
	for {
		select {
		case <-quit:
			fmt.Println("stopped in goroutine")
			return
		default:
			fmt.Println("working, now wait for 1 second in goroutine")
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now())
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
	fmt.Println(fmt.Sprintf("main goroutine finished: %s", time.Now()))
}
