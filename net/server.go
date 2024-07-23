package main

//http server

import (
	"fmt"
	"net/http"
)

// handler function

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
