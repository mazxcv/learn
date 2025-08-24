package pangram

import (
	"regexp"
	"strings"
)

type Pangram map[string]int

func makePangram() map[string]int {
	return Pangram{}
}

// IsPangram checks if a phrase is a pangram
func IsPangram(input string) bool {
	isWord := regexp.MustCompile(`[a-zA-Z]`)
	pangram := makePangram()
	for _, v := range strings.Split(strings.ToLower(input), "") {
		if isWord.MatchString(v) {
			pangram[v] = 1
		}
	}

	keys := make([]string, 0, len(pangram))

	for k := range pangram {
		keys = append(keys, k)
	}
	return len(keys) == 26
}

// // IsPangram checks if a phrase is a pangram
// func IsPangram(s string) bool {
// 	lookup := strings.ToLower(s)
// 	for chr := 'a'; chr <= 'z'; chr++ {
// 		if !strings.ContainsRune(lookup, chr) {
// 			return false
// 		}
// 	}
// 	return true
// }
