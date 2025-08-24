package _8__Length_of_Last_Word

func lengthOfLastWord(s string) (cnt int) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			cnt++
		} else {
			if cnt > 0 {
				break
			}
		}
	}
	return
}
