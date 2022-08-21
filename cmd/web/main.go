package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// mux = http.NewServeMUx: For the sake of security, it's generally a good idea to avoid DefaultServeMux.
	// Use your own locally-scoped servemux instead, like the code below:
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetWrite)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	fmt.Println("Hello World!")
}
