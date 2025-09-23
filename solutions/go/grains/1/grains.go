package grains

import (
    "math"
    "errors"
)

func Square(number int) (uint64, error) {
    if number <= 0 || number > 64 {
        return 0, errors.New("number out of bound")
    }
	return uint64(math.Pow(2.0, float64(number-1))), nil
}

func Total() uint64 {
    var num uint64
	for i := 0;i < 65; i++ {
        numSquare, _ := Square(i)
        num += numSquare
    }
    return num
}
