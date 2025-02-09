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


// TestGetNumberPropertiesHandler tests the /api/classify-number endpoint
func TestGetNumberPropertiesHandler(t *testing.T) {
	// Test cases with different types of numbers
	testCases := []struct {
			name           string
			number         string
			expectedStatus int
			validateFunc   func(*testing.T, *handlers.NumberClassificationResponse)
	}{
			{
					name:           "Valid Armstrong Number",
					number:         "371",
					expectedStatus: http.StatusOK,
					validateFunc: func(t *testing.T, resp *handlers.NumberClassificationResponse) {
							if !resp.IsPrime {
									t.Error("Expected 371 to have prime property")
							}
							if !contains(resp.Properties, "armstrong") {
									t.Error("Expected 371 to be an Armstrong number")
							}
							if !contains(resp.Properties, "odd") {
									t.Error("Expected 371 to be odd")
							}
					},
			},
			{
					name:           "Even Number",
					number:         "100",
					expectedStatus: http.StatusOK,
					validateFunc: func(t *testing.T, resp *handlers.NumberClassificationResponse) {
							if contains(resp.Properties, "odd") {
									t.Error("Expected 100 to be even")
							}
							if contains(resp.Properties, "armstrong") {
									t.Error("Expected 100 to not be an Armstrong number")
							}
					},
			},
			{
					name:           "Invalid Input",
					number:         "abc",
					expectedStatus: http.StatusBadRequest,
					validateFunc: func(t *testing.T, resp *handlers.NumberClassificationResponse) {
							if !resp.Error {
									t.Error("Expected error for invalid input")
							}
					},
			},
	}

	for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
					// Create a request to pass to our handler
					req, err := http.NewRequest("GET", "/api/classify-number?number="+tc.number, nil)
					if err != nil {
							t.Fatal(err)
					}

					// Create a ResponseRecorder to record the response
					rr := httptest.NewRecorder()
					handler := http.HandlerFunc(handlers.GetNumberProperties)

					// Call the handler
					handler.ServeHTTP(rr, req)

					// Check the status code
					if status := rr.Code; status != tc.expectedStatus {
							t.Errorf("handler returned wrong status code: got %v want %v", 
									status, tc.expectedStatus)
					}

					// Only attempt to decode if expecting a successful response
					if tc.expectedStatus == http.StatusOK {
							var response handlers.NumberClassificationResponse
							err = json.Unmarshal(rr.Body.Bytes(), &response)
							if err != nil {
									t.Errorf("failed to decode response: %v", err)
							}

							// Run validation function for the specific test case
							tc.validateFunc(t, &response)
					}
			})
	}
}

// Helper function to check if a slice contains a specific string
func contains(slice []string, str string) bool {
	for _, v := range slice {
			if v == str {
					return true
			}
	}
	return false
}

// Benchmark the number properties handler
func BenchmarkGetNumberProperties(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/classify-number?number=371", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetNumberProperties)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
			handler.ServeHTTP(rr, req)
	}
}