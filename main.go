package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/stephennwachukwu/hng/internal/handlers"
    "github.com/stephennwachukwu/hng/internal/middleware"
)

func main() {
    http.HandleFunc("/", middleware.EnableCORS(handlers.GetInfo))
    http.HandleFunc("/api/classify-number", middleware.EnableCORS(handlers.GetNumberProperties))

    port := ":8080"
    fmt.Printf("Server starting on port %s...\n", port)
    
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
