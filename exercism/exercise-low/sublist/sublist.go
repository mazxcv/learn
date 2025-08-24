package sublist

import (
	"strconv"
	"strings"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	var sb1, sb2 strings.Builder

	for _, v := range l1 {
		sb1.WriteString(strconv.Itoa(v))
		sb1.WriteString(",")
	}
	for _, v := range l2 {
		sb2.WriteString(strconv.Itoa(v))
		sb2.WriteString(",")
	}
	s1 := sb1.String()
	s2 := sb2.String()

	switch {
	case s1 == s2:
		return RelationEqual
	case strings.Contains(s1, s2):
		return RelationSuperlist
	case strings.Contains(s2, s1):
		return RelationSublist
	default:
		return RelationUnequal
	}

}
