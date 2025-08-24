package romannumerals

import "fmt"

type RomanNumeral struct {
	dec   int
	roman string
}

func ConfigDecToRoman() []RomanNumeral {
	return []RomanNumeral{
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
}

func GenerateRomanNumeral(input int) (roman string) {
	for _, v := range ConfigDecToRoman() {
		for input >= v.dec {
			roman += v.roman
			input -= v.dec
		}
	}
	return
}

// func ToRomanNumeral get integer and return romanian view
func ToRomanNumeral(input int) (roman string, err error) {
	if input > 3999 || input <= 0 {
		err = fmt.Errorf("out of range")
		return
	}

	return GenerateRomanNumeral(input), nil
}
