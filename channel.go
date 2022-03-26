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

func main() {
	tooLate := make(chan struct{})
	proCh := make(chan string)

	go func() {
		for {
			fmt.Println("working")
			time.Sleep(1 * time.Second)
			select {
			case <-tooLate:
				fmt.Println("stopped")
				return
			case proCh <- "processed": //this why it won't block the goroutine if the timer expirerd.
			default: // adding default will make it not block
			}
			fmt.Println("done here")

		}
	}()
	select {
	case proc := <-proCh:
		fmt.Println(proc)
	case <-time.After(1 * time.Second):
		fmt.Println("too late")
		close(tooLate)
	}

	time.Sleep(4 * time.Second)
	fmt.Println("finish\n")
}
