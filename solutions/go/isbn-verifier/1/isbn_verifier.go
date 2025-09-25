package isbn

import "unicode"

// runeToInt computes the ISBN-10 checksum.
// It returns (sum, ok) where ok=false if the string is invalid.
func runeToInt(isbn string) (int, bool) {
	sum, cnt := 0, 10

	for _, c := range isbn {
		if c == '-' { // ignore hyphens
			continue
		}

		if cnt == 1 { // last character
			if c == 'X' {
				sum += 10
			} else if unicode.IsDigit(c) {
				sum += int(c - '0')
			} else {
				return 0, false
			}
			cnt--
		} else {
			if unicode.IsDigit(c) {
				sum += int(c-'0') * cnt
				cnt--
			} else {
				return 0, false
			}
		}
	}

	// cnt must end at 0, otherwise string length/format is wrong
	if cnt != 0 {
		return 0, false
	}
	return sum, true
}

func IsValidISBN(isbn string) bool {
	sum, ok := runeToInt(isbn)
	if !ok {
		return false
	}
	return sum%11 == 0
}
