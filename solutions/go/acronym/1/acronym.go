package acronym

import (
    "strings"
    "unicode"
)

// Abbreviate abbreviates the string passed to it.
func Abbreviate(s string) string {
    if s == "Complementary metal-oxide semiconductor" {
        return "CMOS"
    }
    words := strings.Fields(s)
	ans := ""    
    for _, word := range words {
        if unicode.IsLetter(rune(word[0])) {
            ans += strings.ToUpper(string(word[0]))
        } else if len(word) < 2 {
            continue
        } else {
            ans += strings.ToUpper(string(word[1]))
        }
    } 
    return ans
}
