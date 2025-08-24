package _59__Repeated_Substring_Pattern

func repeatedSubstringPattern(s string) bool {

	n := len(s)
	for i := n / 2; i >= 1; i-- {
		candidate := s[:i]
		str := candidate
		if n%i == 0 {
			for j := 1; j < n/i; j++ {
				str += candidate
			}
			if str == s {
				return true
			}
		}
	}
	return false
}
