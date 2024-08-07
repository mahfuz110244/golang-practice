package main

import (
	"fmt"
	"time"
)

func main() {
	dateString := "2023/06/30"
	date, error := time.Parse("2006/01/02", dateString)
	fmt.Println(date)

	dateString = "6/30/2023"
	date, error = time.Parse("1/02/2006", dateString)

	fmt.Println(date)

	fmt.Println(date.AddDate(0, 0, int(10)))
	fmt.Println(time.Now())
	fmt.Println(date.AddDate(0, 0, int(10)).Before(time.Now().UTC().Add(6 * time.Hour)))

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Printf("Type of dateString: %T\n", dateString)
	fmt.Printf("Type of date: %T\n", date)
	fmt.Println()
	fmt.Printf("Value of dateString: %v\n", dateString)
	fmt.Printf("Value of date: %v", date)
	fmt.Printf(time.Now().Format("2006-01-02T15:04:05.000Z"))
}
