// package acronym - is a library to generate acronim
package acronym

import (
	"strings"
)

// Abbreviate return acronym
func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", "")
	var sb strings.Builder
	for _, v := range strings.Fields(s) {

		sb.WriteByte(v[0])
	}
	return strings.ToUpper(sb.String())
}
