package romannumerals

import (
    "strings"
    "errors"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
    if input < 1 || input > 3999 {
        return "", errors.New("invalid input")
    }
    var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for input >= int(numeral.Value) {
			result.WriteString(numeral.Symbol)
			input -= int(numeral.Value)
		}
	}

	return result.String(), nil
}
