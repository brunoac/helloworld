package main

import (
	"fmt"
	"net/http"
)

const defaultPort = ":8080"

var requests int

func main() {
	requests = 0
	fmt.Printf("Started server at localhost:%s\n", defaultPort)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests++
		fmt.Printf("Number of requests: %d\n", requests)
		fmt.Fprint(w, "Hello World!")
	})
	http.ListenAndServe(defaultPort, nil)
}
