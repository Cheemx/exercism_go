package thefarm

import (
    "errors"
    "fmt"
)

type InvalidCowsError struct {
    numOfCows int
    message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numOfCows, e.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numOfCows int) (float64, error) {
    totalFodder, err := fc.FodderAmount(numOfCows)
    if err != nil {
        return 0.0, err
    }
    multiplier, err := fc.FatteningFactor()
    if err != nil {
        return 0.0, err
    }
    return (totalFodder/float64(numOfCows)) * multiplier, nil  
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numOfCows int) (float64, error) {
    if numOfCows > 0 {
        return DivideFood(fc, numOfCows)
    }
    return 0.0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numOfCows int) error {
    if numOfCows < 0 {
        return &InvalidCowsError{
            numOfCows: numOfCows,
            message: "there are no negative cows",
        }
    } else if numOfCows == 0 {
        return &InvalidCowsError{
            numOfCows: numOfCows,
            message: "no cows don't need food",
        }
    }
    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
