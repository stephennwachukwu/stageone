package utils

import (
    "encoding/json"
    "fmt"
    "math"
    "net/http"
    "strconv"
)

func IsPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func IsPerfect(n int) bool {
    sum := 0
    for i := 1; i < n; i++ {
        if n%i == 0 {
            sum += i
        }
    }
    return sum == n
}

func IsArmstrong(n int) bool {
    str := strconv.Itoa(n)
    power := len(str)
    sum := 0
    for _, digit := range str {
        d, _ := strconv.Atoi(string(digit))
        sum += int(math.Pow(float64(d), float64(power)))
    }
    return sum == n
}

func DigitSum(n int) int {
    // Convert negative number to positive
    n = int(math.Abs(float64(n)))
    
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

type NumbersAPIResponse struct {
    Text string `json:"text"`
}

// GetFunFact fetches a math-related fun fact about a number from the Numbers API
func GetFunFact(number int) (string, error) {
    url := fmt.Sprintf("http://numbersapi.com/%d/math?json", number)
    
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("API request failed with status: %d", resp.StatusCode)
    }

    var apiResponse NumbersAPIResponse
    if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
        return "", err
    }

    return apiResponse.Text, nil
}