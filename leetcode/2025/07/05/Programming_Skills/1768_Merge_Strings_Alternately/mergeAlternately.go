package _768_Merge_Strings_Alternately

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	var i int
	for i = 0; i < len(word1) && i < len(word2); i++ {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}

	sb.WriteString(word1[i:])
	sb.WriteString(word2[i:])

	return sb.String()
}
