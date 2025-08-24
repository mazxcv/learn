package cipher

import (
	"strings"
)

// type shift describes shift over a given alphabet using Caesar's algorithm
// Distance = 3 (shift alphabet = 3) (A->D, B->E ...Z->C)
type shift struct {
	Distance int
}

// maybe ddddddddddddd -> 3333333333 shifts
type vigenere struct {
	Distance []int
	key      string
}

func ShiftRune(c rune, distance int) rune {
	var shift int = (int(c-'a') + distance) % 26
	return 'a' + rune(shift)
}

// func NewCaesar - default Caesar cipher (Alphabet 'a'...'z') shift = 3
func NewCaesar() Cipher {
	return &shift{3}
}

func NewShift(distance int) Cipher {
	if distance <= -25 || distance >= 26 || distance == 0 {
		return nil
	}
	if distance > -25 && distance <= 0 {
		distance += 26
	}
	return &shift{distance}
}

func (c shift) Encode(input string) string {
	var sb strings.Builder
	for _, v := range strings.ToLower(input) {
		if v >= 'a' && v <= 'z' {
			sb.WriteRune(ShiftRune(v, c.Distance))
		}
	}
	return sb.String()
}

func (c shift) Decode(input string) string {
	var sb strings.Builder
	for _, v := range strings.ToLower(input) {
		if v >= 'a' && v <= 'z' {
			sb.WriteRune(ShiftRune(v, 26-c.Distance))
		}
	}
	return sb.String()
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	checksum := 0
	distance := []int{}
	for _, v := range key {
		if v >= 'a' && v <= 'z' {
			distance = append(distance, int(v-'a'))
			checksum += int(v - 'a')
		} else {
			return nil
		}
	}
	if checksum == 0 {
		return nil
	}
	return &vigenere{distance, key}
}

func (c vigenere) Encode(input string) string {

	var sb strings.Builder
	i := 0
	for _, v := range strings.ToLower(input) {
		if v >= 'a' && v <= 'z' {
			if i == len(c.Distance) {
				i = 0
			}
			sb.WriteRune(ShiftRune(v, c.Distance[i]))
			i++
		}
	}
	return sb.String()
}

func (c vigenere) Decode(input string) string {
	var sb strings.Builder
	i := 0
	for _, v := range strings.ToLower(input) {
		if v >= 'a' && v <= 'z' {
			if i == len(c.Distance) {
				i = 0
			}
			sb.WriteRune(ShiftRune(v, 26-c.Distance[i]))
			i++
		}
	}
	return sb.String()
}
