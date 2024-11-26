package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Example of using an environment variable for a greeting message
	greeting := os.Getenv("GREETING_MESSAGE")
	if greeting == "" {
		greeting = "Welcome to My Dummy Website"
	}

	fmt.Fprintf(w, "<h1>%s</h1><p>This is a test website served by Go!</p>", greeting)
}

func main() {
	// Use environment variable for port, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	fmt.Printf("Server started at http://127.0.0.1:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
