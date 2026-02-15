package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Register the handler function for the / route
	http.HandleFunc("/", handler)

	// Start the web server on port 3000
	fmt.Println("Starting server on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handler is the function that writes the "Hello, World!" response
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World from go!")
}
