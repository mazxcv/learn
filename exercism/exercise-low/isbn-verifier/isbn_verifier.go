package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	acc := 0
	m := 10

	for _, v := range strings.Split(isbn, "") {
		x, err := strconv.Atoi(v)
		if err != nil {
			switch {
			case v == "X" && m > 1:
				return false
			case v == "X":
				x = 10
			case v == "-":
				continue
			default:
				return false
			}
		}
		acc += m * x
		m--
	}

	return m == 0 && acc%11 == 0
}
