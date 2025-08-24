package _42_Valid_Anagram

func isAnagram(s string, t string) bool {
	tr := map[rune]int{}
	if len(s) != len(t) {
		return false
	}
	for _, v := range s {
		tr[v]++
	}
	for _, v := range t {
		if val, ok := tr[v]; val != 0 && ok {
			tr[v]--
		} else {
			return false
		}
	}
	return true
}
