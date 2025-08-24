package anagram

import "strings"

// func GetSymbolsCount prepare hashMap symbol: count
func GetSymbolsCount(candidate string) map[string]int {
	h := make(map[string]int)
	for _, v := range strings.Split(strings.ToLower(candidate), "") {
		h[v]++
	}
	return h
}

// func EqSymbolCCount return equals 2 hashMap
func EqSymbolCCount(h1, h2 map[string]int) bool {
	if len(h1) != len(h2) {
		return false
	}
	for k, v := range h1 {
		if v != h2[k] {
			return false
		}
	}
	return true
}

// func Detect - return list of anagram subject
func Detect(subject string, candidates []string) (anagram []string) {
	hSub := GetSymbolsCount(strings.ToLower(subject))
	for _, v := range candidates {
		if len(v) != len(subject) {
			continue
		}
		if strings.EqualFold(v, subject) {
			continue
		}
		hCan := GetSymbolsCount(strings.ToLower(v))
		if EqSymbolCCount(hSub, hCan) {
			anagram = append(anagram, v)
		}
	}
	return
}
