package main

import (
	"fmt"
	"net/http"

	"github.com/petermattis/goid" // Import the goid package
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from goroutine %d\n", goid.Get()) // Get current goroutine ID
	fmt.Printf("Hello from goroutine %d\n", goid.Get())     // Get current goroutine ID
}

func main() {
	http.HandleFunc("/", handleRequest)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
