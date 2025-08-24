package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {

	if pt == "" {
		return ""
	}

	pt = strings.Join(strings.FieldsFunc(strings.ToLower(pt), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}), "")

	dx := int(math.Sqrt(float64(len(pt))))
	dy := dx
	switch {
	case dx*dx == len(pt):
		dy = dx
	case dx*(dx+1) >= len(pt):
		dx = dx + 1
	case (dx+1)*(dx+1) >= len(pt):
		dx++
		dy = dx
	}

	arrayRaw := []string{}
	for i := 0; i < dy; i++ {
		var sb strings.Builder
		sb.Grow(dx)
		for j := 0; j < dx; j++ {
			index := i*dx + j
			var symbol byte
			if index >= len(pt) {
				symbol = ' '
			} else {
				symbol = pt[index]
			}
			sb.WriteByte(symbol)
		}
		arrayRaw = append(arrayRaw, sb.String())
	}

	var sb strings.Builder
	sb.Grow(dx*dy + dy - 1)
	for i := 0; i < dx; i++ {

		for j := 0; j < dy; j++ {
			sb.WriteByte(arrayRaw[j][i])
		}
		if i < dx-1 {
			sb.WriteByte(' ')
		}
	}

	return sb.String()
}
