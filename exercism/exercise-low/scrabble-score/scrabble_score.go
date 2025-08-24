// Package scrabble is a library for scores a scrabble word.
package scrabble

import "strings"

// Score computes Scrable score for accepted word.
// Score takes a word and returns its scrabble score.
func Score(word string) (sum int) {
	for _, v := range strings.Split(strings.ToUpper(word), "") {
		switch v {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			sum++
		case "D", "G":
			sum += 2
		case "B", "C", "M", "P":
			sum += 3
		case "F", "H", "V", "W", "Y":
			sum += 4
		case "K":
			sum += 5
		case "J", "X":
			sum += 8
		case "Q", "Z":
			sum += 10
		}
	}
	return
}
