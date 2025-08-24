package validparentheses

type stackParenthnes []rune

func isValid(s string) bool {
	if s == "" {
		return true
	}

	var stack stackParenthnes = make(stackParenthnes, 0)

	for _, v := range s {
		stack.push(v)
	}

	return len(stack) == 0
}

func (s *stackParenthnes) push(r rune) {

	m := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	switch {
	case len(*s) == 0:
		*s = append(*s, r)
	case m[r] == (*s)[len(*s)-1]:
		*s = (*s)[:len(*s)-1]
	default:
		*s = append(*s, r)
	}

}
