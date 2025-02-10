package handlers

import (
    "net/http"
    "time"

    "github.com/stephennwachukwu/hng/internal/utils"
)

type InfoResponse struct {
    Email           string    `json:"email"`
    CurrentDateTime time.Time `json:"current_datetime"`
    GithubURL       string    `json:"github_url"`
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    response := InfoResponse{
        Email:           "stephennwac0007@gmail.com",
        CurrentDateTime: time.Now().UTC(),
        GithubURL:       "https://github.com/stephennwachukwu/stageone",
    }

    utils.WriteJSON(w, http.StatusOK, response)
}