package main

import "fmt"

func main() {
	array := []int{}
	reverseArray := []int{}
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Println(array[i])
		reverseArray = append(reverseArray, array[i])

	}
	fmt.Println(reverseArray)
}
