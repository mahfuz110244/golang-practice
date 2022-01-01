package main

import (
	"fmt"
	"strings"
)

func stringifyUserIDs(IDs []string) string {
	return "[\"" + strings.Join(IDs, "\",\"") + "\"]"
}

func stringifyUserIDsRefactor(IDs []string) string {
	startValue := "[\""
	endValue := "\"]"
	finalValue := startValue + strings.Join(IDs, "\",\"") + endValue
	return finalValue
	return "[\"" + strings.Join(IDs, "\",\"") + "\"]"
}


// Main Method
func main() {
	ids := []string {"user1", "user2", "user3"}
	fmt.Println(stringifyUserIDs(ids))
	fmt.Println(stringifyUserIDsRefactor(ids))
}
