package main

import (
	"fmt"
	"strings"
	"sync"
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

func Run(tasks ...func(*sync.WaitGroup, chan<- error)) error {
	n := len(tasks)
	errCh := make(chan error, n)

	w := &sync.WaitGroup{}
	w.Add(n)

	for _, t := range tasks {
		go t(w, errCh)
	}

	w.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			return err
		}
	}
	return nil
}



// Main Method
func main() {
	ids := []string {"user1", "user2", "user3"}
	fmt.Println(stringifyUserIDs(ids))
	fmt.Println(stringifyUserIDsRefactor(ids))
}
