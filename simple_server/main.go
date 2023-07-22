package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a route handler for the root ("/") URL path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") // Send "Hello, World!" as the response
	})

	// Serve static files from the "static" directory
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}