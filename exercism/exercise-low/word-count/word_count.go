package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

// func WordCount counts all words in phrase and return
func WordCount(phrase string) Frequency {
	f := make(Frequency)

	r := regexp.MustCompile(`([\w\d]+[\w\d\']*[\w\d]+)|([\w\d]+)`)

	for _, v := range r.FindAllString(strings.Trim(phrase, "\n\t"), -1) {
		f[strings.ToLower(v)]++
	}

	return f
}
