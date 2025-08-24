package _8_Find_the_Index_of_the_First_Occurrence_in_a_String

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 && len(needle) != 0 {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if len(needle)+i > len(haystack) {
				return -1
			}
			flagOccurence := true
			for j := 1; j < len(needle) && flagOccurence; j++ {
				flagOccurence = haystack[i+j] == needle[j]
			}
			if flagOccurence {
				return i
			}
		}
	}
	return -1
}
