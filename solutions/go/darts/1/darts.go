package darts

import "math"

func Score(x, y float64) int {
	sqrt := math.Sqrt(x*x + y*y)
    if sqrt <= 1.0 {
        return 10
    } else if sqrt <= 5.0 {
        return 5
    } else if sqrt <= 10.0 {
        return 1
    }
    return 0
}
