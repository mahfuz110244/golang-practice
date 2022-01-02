package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// func stringifyUserIDs(IDs []string) string {
// 	return "[\"" + strings.Join(IDs, "\",\"") + "\"]"
// }

func stringifyUserIDs(IDs []string) string {
	startStr := "[\""
	endStr := "\"]"
	finalStr := startStr + strings.Join(IDs, "\",\"") + endStr
	return finalStr
}

//a.) return multiple errors,
//b.) support task cancellation.

// func Run(tasks ...func(*sync.WaitGroup, chan<- error)) error {
// 	n := len(tasks)
// 	errCh := make(chan error, n)

// 	w := &sync.WaitGroup{}
// 	w.Add(n)

// 	for _, t := range tasks {
// 		go t(w, errCh)
// 	}

// 	w.Wait()
// 	close(errCh)
// 	for err := range errCh {
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func Run(tasks ...func(*sync.WaitGroup, chan<- error)) []error {
	n := len(tasks)
	errCh := make(chan error, n)

	w := &sync.WaitGroup{}
	w.Add(n)

	//Make a background context
	ctx := context.Background()

	for _, t := range tasks {
		_, cancelFunction := context.WithTimeout(ctx, 1*time.Second)
		defer cancelFunction()

		go t(w, errCh)
	}

	w.Wait()
	close(errCh)
	errors := []error{}
	for err := range errCh {
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}

// Main Method
func main() {
	ids := []string{"user1", "user2", "user3"}
	fmt.Println(stringifyUserIDs(ids))
	// fmt.Println(stringifyUserIDsRefactor(ids))
}
