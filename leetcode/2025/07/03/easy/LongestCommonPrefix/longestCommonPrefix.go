package longestcommonprefix

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {

	sort.Strings(strs)

	var i int
	for i < len(strs[0]) {
		if strs[0][i] != strs[len(strs)-1][i] {
			break
		}
		i++
	}

	return strs[0][:i]
}
