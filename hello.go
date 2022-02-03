package main

import "fmt"

func HelloWorld(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

func main() {
	fmt.Println(HelloWorld("World"))
}
