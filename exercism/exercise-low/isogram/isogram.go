// Package isogram is a small library for analyzing if a phrase is an isogram.
package isogram

import (
	"strings"
)

// IsIsogram accepts a word and returns if it is an isogram.
func IsIsogram(word string) bool {
	testIsogram := map[string]int{}
	for _, v := range strings.Split(strings.ToLower(word), "") {
		if _, ok := testIsogram[v]; ok {
			if v != " " && v != "-" {
				return false
			}
		}
		testIsogram[v]++
	}
	return true
}
