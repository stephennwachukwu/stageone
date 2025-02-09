package utils

import (
    "math"
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
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}