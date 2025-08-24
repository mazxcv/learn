package luhn

import (
	"strconv"
	"strings"
)

// func Valid determine id-card is valid per luhn-formula
func Valid(id string) bool {
	digits := []int{}

	for _, v := range strings.Split(id, "") {
		if v == " " {
			continue
		}
		v, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		digits = append(digits, v)
	}
	len := len(digits)
	if len <= 1 {
		return false
	}
	sum := 0

	for i, v := range digits {
		if (i+len)%2 == 0 {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		sum += v
	}
	return sum%10 == 0
}
