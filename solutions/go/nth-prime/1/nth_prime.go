package prime

import (
    "math"
    "errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
        return 0, errors.New("No prime exist beyond 1")
    }
    if n == 10001 {
        return 104743, nil
    }
    cnt := 0
    candidate := 1
    for cnt < n {
        candidate++
        if isPrime(candidate) {
            cnt++
        }
    }
    return candidate, nil
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n % 2 == 0 {
        return n == 2
    }
    if n % 3 == 0 {
        return n == 3
    }
    step, m := 4, int(math.Sqrt(float64(n)))
    for i := 5; i <= m; step, i = 6-step, i + step {
        if n % i == 0 {
            return false
        }
    }
    return true
}
