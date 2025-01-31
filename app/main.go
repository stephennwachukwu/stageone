package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

// InfoResponse represents the required response structure
type InfoResponse struct {
    Email           string    `json:"email"`
    CurrentDateTime time.Time `json:"current_datetime"`
    GithubURL      string    `json:"github_url"`
}

// MarshalJSON implements custom JSON marshaling to format the timestamp
func (i InfoResponse) MarshalJSON() ([]byte, error) {
    type Alias InfoResponse
    return json.Marshal(&struct {
        Email           string `json:"email"`
        CurrentDateTime string `json:"current_datetime"`
        GithubURL      string `json:"github_url"`
    }{
        Email:           i.Email,
        CurrentDateTime: i.CurrentDateTime.Format(time.RFC3339),
        GithubURL:      i.GithubURL,
    })
}

// enableCORS is a middleware that handles CORS headers
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // Handle preflight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next(w, r)
    }
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    
    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.Printf("Error encoding JSON: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}

func getInfo(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    response := InfoResponse{
        Email:           "stephennwac007@gmail.com",
        CurrentDateTime: time.Now().UTC(),
        GithubURL:      "https://github.com/stephennwachukwu/stageone",
    }

    writeJSON(w, http.StatusOK, response)
}

func main() {
    http.HandleFunc("/", enableCORS(getInfo))

    port := ":8087"
    fmt.Printf("Server starting on port %s...\n", port)
    
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}