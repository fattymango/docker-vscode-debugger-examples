package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the server...")
	// simple http server with a single endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		// Hello, World!
		fmt.Fprintf(w, "Hello, World!")
	})

	// start the server
	http.ListenAndServe("0.0.0.0:8080", nil)

}
