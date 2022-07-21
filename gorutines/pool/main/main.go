package main

import (
	"fmt"
	"strconv"
	"sync"
)

func worker(input chan string, output chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range input {
		output <- value + " processed"
	}
}

func main() {
	wg := &sync.WaitGroup{}
	var jobs = []string{}
	for i := 1; i <= 50; i++ {
		jobs = append(jobs, "Job "+strconv.Itoa(i))
	}
	fmt.Println(jobs)

	input := make(chan string, len(jobs))
	output := make(chan string, len(jobs))

	numOfWorkers := 10
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(input, output, wg)
	}

	for _, job := range jobs {
		input <- job
	}
	close(input)

	wg.Wait()

	close(output)

	for result := range output {
		fmt.Println(result)
	}

}
