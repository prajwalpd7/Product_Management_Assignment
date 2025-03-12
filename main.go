package main

import (
    "fmt"
    "net/http"
    "time"
)

// timeHandler writes the current date & time to the response.
func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("2006-01-02 15:04:05 MST")
    fmt.Fprintf(w, "Current Date & Time: %s", currentTime)
}


func main() {
    http.HandleFunc("/", timeHandler)
    fmt.Println("Server is starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}