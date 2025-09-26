package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
    plainText := "abcdefghijklmnopqrstuvwxyz"
	cipher := plainText[shiftKey:] + plainText[:shiftKey]
    if shiftKey == 0 || shiftKey == 26 {
        return plain
    }

    mepP := make(map[rune]int)
    for i, c := range plainText {
        mepP[c] = i
    }
    
    str := ""
    for _, c := range plain {
        if unicode.IsUpper(c) {
            char := unicode.ToLower(c)
            str += string(unicode.ToUpper(rune(cipher[mepP[char]])))
        } else if !unicode.IsLetter(c){
            str += string(c)
        } else{
        	str += string(cipher[mepP[c]])
        }
    }
    return str
}
