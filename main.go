package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

// Default values, will be overridden by build flags
var (
    GREETING_MESSAGE = "Default Greeting"
    PORT             = "8080"
)

func main() {
    // Use environment variables if they exist, otherwise use compiled defaults
    greeting := os.Getenv("GREETING_MESSAGE")
    if greeting == "" {
        greeting = GREETING_MESSAGE
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = PORT
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, greeting)
    })

    log.Printf("Starting server on port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
