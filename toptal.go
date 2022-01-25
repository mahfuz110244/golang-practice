package main

import (
	"fmt"
	s "strings"
)

func GetCropMessage(message string, k int) string {
	array := s.Split(message, " ")
	messageArray := s.Split(message[:k], " ")
	lastMessage := messageArray[len(messageArray)-1]
	found := false
	for _, v := range array {
		// fmt.Println(v, lastMessage)
		if v == lastMessage {
			found = true
		}
	}
	// fmt.Println(found)
	if !found {
		messageArray = messageArray[:len(messageArray)-1]
	}
	fmt.Println(s.Join(messageArray, " "))
	return s.Join(messageArray, " ")
}

func main() {
	// Solution 1
	message := "Codibility We test coders"
	k := 17
	GetCropMessage(message, k)
}
