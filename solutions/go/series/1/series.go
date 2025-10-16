package series

func All(n int, s string) []string {
	words := []string{}
    for i := 0;i <= len(s)-n; i++ {
        word := s[i:i+n]
        words = append(words, word)
    }
    return words
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}
