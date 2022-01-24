package main

import (
	"fmt"
	"strings"
)

// Main Method
func main() {
	str1 := "custom_error!!!calculation error!!! sales order item given total = 2000 missmatch with items calculated total = 200"
	str1 = "custom_error!!!"

	res1 := strings.Split(str1, "custom_error!!!")

	// Displaying the result

	fmt.Println("\nResult 1: ", len(res1))
	fmt.Printf("type of a is %T\n", res1)
}
