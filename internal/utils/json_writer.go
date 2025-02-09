package utils

import (
    "encoding/json"
    "log"
    "net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    
    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.Printf("Error encoding JSON: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}