package cipher

import (
    "strings"
    "unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
    distance int
}

type vigenere struct {
    keyIndices []int
}

func NewCaesar() Cipher {
    return shift{
        distance: 3,
    }
}

func NewShift(distance int) Cipher {
    if distance <= -26 || distance >= 26 || distance == 0 {
        return nil
    }
    return shift{
        distance: distance,
    }
}

func (c shift) Encode(input string) string {
    var result strings.Builder
    input = strings.ToLower(input)
    for _, ch := range input {
        if unicode.IsLetter(ch) {
            base := int('a')
            shifted := (int(ch)-base + c.distance + 26) % 26
            result.WriteRune(rune(base + shifted))
        }
    }
    return result.String()
}

func (c shift) Decode(input string) string {
    var result strings.Builder
    input = strings.ToLower(input)
    for _, ch := range input {
        if unicode.IsLetter(ch) {
            base := int('a')
            shifted := (int(ch)-base - c.distance + 26) % 26
            result.WriteRune(rune(base + shifted))
        }
    }
    return result.String()
}

func NewVigenere(key string) Cipher {
    if len(key) == 0 {
        return nil
    }

    keyIndices := []int{}
    allA := true

    for _, char := range key {
        if !unicode.IsLower(char) {
            // reject uppercase or non-lowercase letters
            return nil
        }
        idx := int(char - 'a')
        if idx != 0 {
            allA = false
        }
        keyIndices = append(keyIndices, idx)
    }

    if allA {
        return nil
    }

    return vigenere{
        keyIndices: keyIndices,
    }
}


func (v vigenere) Encode(input string) string {
    var result strings.Builder
    input = strings.ToLower(input)
    j := 0

    for _, ch := range input {
        if unicode.IsLetter(ch) {
            base := int('a')
            shift := v.keyIndices[j%len(v.keyIndices)]
            encoded := (int(ch)-base + shift) % 26
            result.WriteRune(rune(base + encoded))
            j++
        }
    }

    return result.String()
}

func (v vigenere) Decode(input string) string {
    var result strings.Builder
    input = strings.ToLower(input)
    j := 0

    for _, ch := range input {
        if unicode.IsLetter(ch) {
            base := int('a')
            shift := v.keyIndices[j%len(v.keyIndices)]
            decoded := (int(ch)-base - shift + 26) % 26
            result.WriteRune(rune(base + decoded))
            j++
        }
    }

    return result.String()
}