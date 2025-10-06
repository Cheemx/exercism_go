package atbash

import (
    "unicode"
    "strings"
)

var mep = map[rune]rune{
    'a': 'z',
    'b': 'y',
    'c': 'x',
    'd': 'w',
    'e': 'v',
    'f': 'u',
    'g': 't',
    'h': 's',
    'i': 'r',
    'j': 'q',
    'k': 'p',
    'l': 'o',
    'm': 'n',
    'n': 'm',
    'o': 'l',
    'p': 'k',
    'q': 'j',
    'r': 'i',
    's': 'h',
    't': 'g',
    'u': 'f',
    'v': 'e',
    'w': 'd',
    'x': 'c',
    'y': 'b',
    'z': 'a',
}

func Atbash(s string) string {
    encoded := ""
    cnt := 0
    for _, char := range strings.ToLower(s) {
        if cnt == 5 {
            encoded += string(" ")
            cnt = 0
        }
        if unicode.IsLetter(char) {
            encoded += string(mep[char])
            cnt++
        } else if unicode.IsDigit(char) {
            encoded += string(char)
            cnt++
        }
    }
    encoded = strings.TrimSpace(encoded)
    return encoded
}
