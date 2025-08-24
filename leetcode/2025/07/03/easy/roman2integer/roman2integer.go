package roman2integer

import "strings"

// основная идея:
// вычитываем XL -
// первый символ (10)
// следующий символ - 50, он больше (10), значит к аккумулятору надо добавить (50) - (10) - (10),
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	roma := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	prev := 2000
	acc := 0
	for _, v := range strings.Split(s, "") {
		if roma[v] > prev {
			acc += roma[v] - prev - prev
		} else {
			acc += roma[v]
		}
		prev = roma[v]
	}

	return acc

}
