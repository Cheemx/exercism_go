package pangram

import "strings"

func IsPangram(input string) bool {
    input = strings.ToLower(input)
    if len(input) < 26 {
        return false
    }
    charBool := map[rune]bool{
        'a': false,
        'b': false,
        'c': false,
        'd': false,
        'e': false,
        'f': false,
        'g': false,
        'h': false,
        'i': false,
        'j': false,
        'k': false,
        'l': false,
        'm': false,
        'n': false,
        'o': false,
        'p': false,
        'q': false,
        'r': false,
        's': false,
        't': false,
        'u': false,
        'v': false,
        'w': false,
        'x': false,
        'y': false,
        'z': false,        
    }

    for _, c := range input {
        charBool[c] = true
    }
    for _, v := range charBool {
        if v == false {
            return false
        }
    }
    return true
}
