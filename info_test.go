package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

		"github.com/stephennwachukwu/hng/internal/handlers"
    "github.com/stephennwachukwu/hng/internal/middleware"
)

type TestResponse struct {
	Email           string `json:"email"`
	CurrentDateTime string `json:"current_datetime"`
	GithubURL      string `json:"github_url"`
}

func TestGetInfo(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	
	// Create a handler wrapped with CORS middleware
	handler := middleware.EnableCORS(handlers.GetInfo)
	
	// Call the handler directly
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the response body
	var response *TestResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Validate response fields
	if response.Email == "" {
		t.Error("Email should not be empty")
	}

	if response.GithubURL == "" {
		t.Error("GitHub URL should not be empty")
	}

	// Parse the current datetime to ensure it's a valid time
	_, err = time.Parse(time.RFC3339, response.CurrentDateTime)
	if err != nil {
		t.Errorf("Invalid datetime format: %v", err)
	}
}

func TestGetInfoMethodNotAllowed(t *testing.T) {
	// Create a POST request (not allowed)
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	
	// Create a handler wrapped with CORS middleware
	handler := middleware.EnableCORS(handlers.GetInfo)
	
	// Call the handler directly
	handler.ServeHTTP(rr, req)

	// Check that we get a Method Not Allowed status
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestCORSHeaders(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	
	// Create a handler wrapped with CORS middleware
	handler := middleware.EnableCORS(handlers.GetInfo)
	
	// Call the handler directly
	handler.ServeHTTP(rr, req)

	// Check CORS headers
	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type, Authorization",
	}

	for header, expectedValue := range expectedHeaders {
		value := rr.Header().Get(header)
		if value != expectedValue {
			t.Errorf("Incorrect %s header: got %v want %v", header, value, expectedValue)
		}
	}
}
