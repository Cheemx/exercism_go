package strain

// Implement the "Keep" and "Discard" function in this file.

func Keep[T any](cp []T, f func(x T) bool) []T {
    var ans []T
    for _, t := range cp {
        if f(t) {
            ans = append(ans, t)
        }
    }
    return ans
}

func Discard[T any](cp []T, f func(x T) bool) []T {
    var ans []T
    for _, t := range cp {
        if !f(t) {
            ans = append(ans, t)
        }
    }
    return ans
}