package anagram

import (
    "unicode"
    "strings"
)

func areMapsEqual(map1, map2 map[rune]int) bool {
    if len(map1) != len(map2) {
        return false
    }

    for key, value := range map1 {
		if val, ok := map2[key]; !ok || val != value {
			return false 
		}
	}
    return true
}

func Detect(subject string, candidates []string) []string {
	runeFreq := map[rune]int{}
    for _, char := range subject {
        lowerChar := char
        if unicode.IsLetter(char) {
            lowerChar = unicode.ToLower(char)
        }
        runeFreq[lowerChar]++
    }

    var ans []string
    for _, word := range candidates {
        if strings.ToLower(word) == strings.ToLower(subject) {
            continue
        }
        wordRuneFreq := map[rune]int{}
        for _, char := range word {
            lowerChar := char
            if unicode.IsLetter(char) {
                lowerChar = unicode.ToLower(char)
            }
            wordRuneFreq[lowerChar]++
        }
        if areMapsEqual(wordRuneFreq, runeFreq) {
        	ans = append(ans, word)
        }
    }
    return ans
}
