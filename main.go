package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client
	// Importantly, we then return from handler. If we don't return the handler
	// would keep executing and also write the "Hello from snippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetWrite(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	//if r.Method != "POST" {
	if r.Method != http.MethodPost {
		// if it's not, use the w.WriteHeader() method to send a 405 status code and w.Write() method to write
		// a "Method Not Allowed" response body. We then return from the function so that the subsequent code is not
		// executed
		// Use the Header().Set() method to add an 'Allow: POST' header to response header map. The first parameter is
		// the header name, and the second parameter is the header vale.
		w.Header().Set("Allow", http.MethodPost)
		// It's only possible to call w.WriteHeader() once per response. If you don't call w.WriteHeader() explicitly,
		// then the first call to w.Write() will automatically send a 200 OK status code to the user.
		//w.WriteHeader(405)
		//w.Write([]byte("Method Not Allowed"))
		// This is a lightweight helper function which takes a given message and status code, then calls the
		// w.WriteHeader() and w.Write() method behind-the-scenes for us
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

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
