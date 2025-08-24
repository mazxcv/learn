package series

func All(n int, s string) (all []string) {
	if n > len(s) {
		return []string{}
	}
	for i := 0; i <= len(s)-n; i++ {
		all = append(all, s[i:i+n])
	}
	return
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
