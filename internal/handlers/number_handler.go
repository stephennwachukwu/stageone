package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"

    "github.com/stephennwachukwu/hng/internal/utils"
)

type NumberClassificationResponse struct {
    Number      interface{} `json:"number"`
    IsPrime     bool        `json:"is_prime"`
    IsPerfect   bool        `json:"is_perfect"`
    Properties  []string    `json:"properties"`
    DigitSum    int         `json:"digit_sum"`
    FunFact     string      `json:"fun_fact"`
    Error       bool        `json:"error,omitempty"`
}

func GetNumberProperties(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    numberParam := r.URL.Query().Get("number")
    
    number, err := strconv.Atoi(numberParam)
    if err != nil {
        response := NumberClassificationResponse{
            Number: numberParam,
            Error:  true,
        }
        utils.WriteJSON(w, http.StatusBadRequest, response)
        return
    }

    properties := []string{}
    if utils.IsArmstrong(number) {
        properties = append(properties, "armstrong")
    }
    if number%2 == 0 {
        properties = append(properties, "even")
    } else {
        properties = append(properties, "odd")
    }

    funFact := fmt.Sprintf("%d is %s", number, 
        strings.Join(properties, " and "))
    if utils.IsArmstrong(number) {
        funFact += fmt.Sprintf(" because the sum of its digits raised to the power of %d equals %d", 
            len(strconv.Itoa(number)), number)
    }

    response := NumberClassificationResponse{
        Number:      number,
        IsPrime:     utils.IsPrime(number),
        IsPerfect:   utils.IsPerfect(number),
        Properties:  properties,
        DigitSum:    utils.DigitSum(number),
        FunFact:     funFact,
    }

    utils.WriteJSON(w, http.StatusOK, response)
}