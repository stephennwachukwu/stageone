package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/stephennwachukwu/hng/internal/utils"
)

type NumberClassificationResponse struct {
    Number    interface{} `json:"number"`
    IsPrime   bool       `json:"is_prime,omitempty"`
    IsPerfect bool       `json:"is_perfect,omitempty"`
    Properties []string  `json:"properties,omitempty"`
    DigitSum  int        `json:"digit_sum,omitempty"`
    FunFact   string     `json:"fun_fact,omitempty"`
    Error     bool       `json:"error,omitempty"`
}

func GetNumberProperties(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    numberParam := r.URL.Query().Get("number")
    
    number, err := strconv.Atoi(numberParam)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(NumberClassificationResponse{
            Number: numberParam,
            Error:  true,
        })
        return
    }

    // Determine properties based on specification
    properties := []string{}
    if utils.IsArmstrong(number) {
        if number%2 == 0 {
            properties = append(properties, "armstrong", "even")
        } else {
            properties = append(properties, "armstrong", "odd")
        }
    } else {
        if number%2 == 0 {
            properties = append(properties, "even")
        } else {
            properties = append(properties, "odd")
        }
    }

    // Get fun fact from Numbers API (math type)
    funFact, err := utils.GetFunFact(number)
    if err != nil || funFact == "" {
        // If Numbers API fails, generate Armstrong number fact if applicable
        if utils.IsArmstrong(number) {
            str := strconv.Itoa(number)
            power := len(str)
            funFact = fmt.Sprintf("%d is an Armstrong number because ", number)
            for i, digit := range str {
                d, _ := strconv.Atoi(string(digit))
                funFact += fmt.Sprintf("%d^%d", d, power)
                if i < len(str)-1 {
                    funFact += " + "
                }
            }
            funFact += fmt.Sprintf(" = %d", number)
        } else {
            funFact = fmt.Sprintf("%d is %s", number, properties[0])
        }
    }

    response := NumberClassificationResponse{
        Number:    number,
        IsPrime:   utils.IsPrime(number),
        IsPerfect: utils.IsPerfect(number),
        Properties: properties,
        DigitSum:  utils.DigitSum(number),
        FunFact:   funFact,
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}