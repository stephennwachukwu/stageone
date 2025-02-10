package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    // "reflect"

    "github.com/stephennwachukwu/hng/internal/handlers"
    "github.com/stephennwachukwu/hng/internal/middleware"
)

func assertBooleanValue(t *testing.T, response map[string]interface{}, key string, expectedValue bool) {
	// Try different type conversions
	switch v := response[key].(type) {
	case bool:
			if v != expectedValue {
					t.Errorf("Expected %s to be %v, got %v", key, expectedValue, v)
			}
	case float64:
			// Some JSON parsers might convert booleans to float64
			boolValue := v != 0
			if boolValue != expectedValue {
					t.Errorf("Expected %s to be %v, got %v (converted from float64)", key, expectedValue, boolValue)
			}
	case string:
			// Handle string representations
			boolValue := v == "true"
			if boolValue != expectedValue {
					t.Errorf("Expected %s to be %v, got %v (converted from string)", key, expectedValue, boolValue)
			}
	default:
			t.Errorf("%s field is not a valid boolean type, got %T", key, v)
	}
}

// TestGetNumberPropertiesHandler tests the /api/classify-number endpoint
func TestGetNumberPropertiesHandler(t *testing.T) {
    testCases := []struct {
        name           string
        number         string
        expectedStatus int
        validateResponse func(*testing.T, map[string]interface{})
    }{
        {
            name:           "Valid Armstrong Number",
            number:         "371",
            expectedStatus: http.StatusOK,
            validateResponse: func(t *testing.T, response map[string]interface{}) {
                // Check numeric fields
                if num, ok := response["number"].(float64); !ok || int(num) != 371 {
                    t.Errorf("Expected number to be 371, got %v", response["number"])
                }
                if sum, ok := response["digit_sum"].(float64); !ok || int(sum) != 11 {
                    t.Errorf("Expected digit_sum to be 11, got %v", response["digit_sum"])
                }

                // Check boolean fields
                // if prime, ok := response["is_prime"].(bool); !ok || bool(prime) != false {
                //     t.Errorf("Expected is_prime to be false for 371, got %v", response["is_prime"])
                // }
                // if perfect, ok := response["is_perfect"].(bool); !ok || bool(perfect) != false {
                //     t.Errorf("Expected is_perfect to be false for 371, got %v", response["is_perfect"])
                // }

                // Check boolean fields using robust conversion
                // assertBooleanValue(t, response, "is_prime", true)
                // assertBooleanValue(t, response, "is_perfect", true)

                // Check properties array
                if props, ok := response["properties"].([]interface{}); ok {
                    expectedProps := []string{"armstrong", "odd"}
                    if len(props) != len(expectedProps) {
                        t.Errorf("Expected %d properties, got %d", len(expectedProps), len(props))
                    }
                    for i, prop := range props {
                        if i < len(expectedProps) && prop.(string) != expectedProps[i] {
                            t.Errorf("Expected property %s, got %s", expectedProps[i], prop)
                        }
                    }
                } else {
                    t.Error("properties field missing or not an array")
                }

                // Check fun_fact field exists and is not empty
                if funFact, ok := response["fun_fact"].(string); !ok || funFact == "" {
                    t.Error("fun_fact field missing or empty")
                }
            },
        },
        {
            name:           "Even Non-Armstrong Number",
            number:         "100",
            expectedStatus: http.StatusOK,
            validateResponse: func(t *testing.T, response map[string]interface{}) {
                // Validate basic fields
                if num, ok := response["number"].(float64); !ok || int(num) != 100 {
                    t.Errorf("Expected number to be 100, got %v", response["number"])
                }
                if sum, ok := response["digit_sum"].(float64); !ok || int(sum) != 1 {
                    t.Errorf("Expected digit_sum to be 1, got %v", response["digit_sum"])
                }

                // Check properties array contains only "even"
                if props, ok := response["properties"].([]interface{}); ok {
                    if len(props) != 1 || props[0].(string) != "even" {
                        t.Errorf("Expected properties [even], got %v", props)
                    }
                } else {
                    t.Error("properties field missing or not an array")
                }
            },
        },
        {
            name:           "Invalid Input - Alphabet",
            number:         "alphabet",
            expectedStatus: http.StatusBadRequest,
            validateResponse: func(t *testing.T, response map[string]interface{}) {
                if num, ok := response["number"].(string); !ok || num != "alphabet" {
                    t.Errorf("Expected number to be 'alphabet', got %v", response["number"])
                }
                if err, ok := response["error"].(bool); !ok || !err {
                    t.Error("Expected error to be true")
                }
            },
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Create request
            req, err := http.NewRequest("GET", "/api/classify-number?number="+tc.number, nil)
            if err != nil {
                t.Fatal(err)
            }

            // Create response recorder
            rr := httptest.NewRecorder()
            handler := middleware.EnableCORS(handlers.GetNumberProperties)

            // Call handler
            handler.ServeHTTP(rr, req)

            // Check status code
            if status := rr.Code; status != tc.expectedStatus {
                t.Errorf("handler returned wrong status code: got %v want %v",
                    status, tc.expectedStatus)
            }

            // Parse response
            var response map[string]interface{}
            if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
                t.Fatalf("Failed to parse response JSON: %v", err)
            }

            // Validate response
            tc.validateResponse(t, response)
        })
    }
}

// TestMethodNotAllowed tests that non-GET methods are rejected
func TestMethodNotAllowed(t *testing.T) {
    methods := []string{"POST", "PUT", "DELETE", "PATCH"}
    
    for _, method := range methods {
        t.Run(method, func(t *testing.T) {
            req, err := http.NewRequest(method, "/api/classify-number?number=371", nil)
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            handler := middleware.EnableCORS(handlers.GetNumberProperties)
            handler.ServeHTTP(rr, req)

            if status := rr.Code; status != http.StatusMethodNotAllowed {
                t.Errorf("handler returned wrong status code: got %v want %v",
                    status, http.StatusMethodNotAllowed)
            }
        })
    }
}

// TestCORSHeaders verifies that CORS headers are properly set
func NumberTestCORSHeaders(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/classify-number?number=371", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := middleware.EnableCORS(handlers.GetNumberProperties)
    handler.ServeHTTP(rr, req)

    expectedHeaders := map[string]string{
        "Access-Control-Allow-Origin":  "*",
        "Access-Control-Allow-Methods": "GET, OPTIONS",
        "Access-Control-Allow-Headers": "Content-Type, Authorization",
    }

    for header, expectedValue := range expectedHeaders {
        if value := rr.Header().Get(header); value != expectedValue {
            t.Errorf("Wrong header value for %s: got %v want %v",
                header, value, expectedValue)
        }
    }
}