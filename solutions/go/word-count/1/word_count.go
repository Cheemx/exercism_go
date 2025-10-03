package wordcount

import (
    "strings"
    "regexp"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	mep := Frequency{}
    phrase = strings.ToLower(phrase)
    re := regexp.MustCompile(`[a-z0-9]+(?:'[a-z0-9]+)?`)

	words := re.FindAllString(phrase, -1)
	for _, word := range words {
		mep[word]++
	}
    return mep
}
