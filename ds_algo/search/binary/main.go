package main

import "fmt"

func binarySearch(array []int, item int) int {
	lower := 0
	upper := len(array)
	for lower < upper {
		mid := (lower + upper) / 2
		if array[mid] == item {
			fmt.Println(mid)
			return mid
			break
		} else {
			if array[mid] < item {
				lower = mid + 1
			} else {
				upper = mid - 1
			}
		}
	}
	return -1
}

func main() {
	// array := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	array := []int{0, 10}
	item := 1000
	fmt.Println(binarySearch(array, item))
}
