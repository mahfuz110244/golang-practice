package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello")
}

const serverPort int = 8000

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Server run port: %d\n", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)

}
