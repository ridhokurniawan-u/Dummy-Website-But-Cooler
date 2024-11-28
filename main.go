package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    greeting := os.Getenv("GREETING_MESSAGE")
    if greeting == "" {
        greeting = "Default Greeting"
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, greeting)
    })

    log.Printf("Starting server on port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
