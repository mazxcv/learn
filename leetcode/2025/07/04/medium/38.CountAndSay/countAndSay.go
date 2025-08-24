package countandsay

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	s := "11"
	for k := 2; k < n; k++ {
		var sb strings.Builder
		sa := strings.Split(s, "")
		c := sa[0]
		count := 1
		for i := 1; i < len(s); i++ {
			if c == sa[i] {
				count++
			} else {
				sb.WriteString(fmt.Sprintf("%d%s", count, c))
				count = 1
				c = sa[i]
			}
		}
		sb.WriteString(fmt.Sprintf("%d%s", count, c))
		s = sb.String()
	}

	return s
}
