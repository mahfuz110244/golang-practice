package main

import (
	"fmt"
	"strconv"
)

// Main Method
func main() {
	outBoolean, _ := strconv.ParseBool("false")
	fmt.Println("String false to boolean converted : ", outBoolean)

	ourInteger, _ := strconv.Atoi("12345")
	fmt.Println("String 12345 to Integer converted : ", ourInteger)

	ourString := strconv.Itoa(12345)
	fmt.Println("Integer 12345 to String converted : ", ourString)

	ourInteger64, _ := strconv.ParseInt("12345", 10, 64)
	fmt.Println("String 12345 to Integer64 converted : ", ourInteger64)

	ourInteger32, _ := strconv.ParseInt("12345", 10, 32)
	fmt.Println("String 12345 to Integer32 converted : ", ourInteger32)

}
