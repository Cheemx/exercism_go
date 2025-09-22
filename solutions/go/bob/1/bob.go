package bob

import (
	"strings"
	"unicode"
)

func IsShouting(s string) bool {
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetter
}


func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	if trimmed == "" {
		return "Fine. Be that way!"
	}

	question := strings.HasSuffix(trimmed, "?")
	shouting := IsShouting(trimmed)

	switch {
	case question && shouting:
		return "Calm down, I know what I'm doing!"
	case shouting:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	default:
		return "Whatever."
	}
}

