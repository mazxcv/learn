package strstr

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 && len(needle) != 0 {
		return -1
	}

	for i := 0; i < len(haystack) && haystack[i] == needle[0]; i++ {
		flagOccurence := true
		for j := 1; j < len(needle) && flagOccurence; j++ {
			flagOccurence = haystack[i+j] == needle[j]
		}
		if flagOccurence {
			return i
		}
	}
	return -1
}
