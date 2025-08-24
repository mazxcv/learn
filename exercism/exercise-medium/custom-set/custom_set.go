package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

type Set map[string]bool

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, v := range l {
		s[v] = true
	}
	return s
}

func (s Set) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	k := 0
	for i := range s {
		k++
		sb.WriteString("\"")
		sb.WriteString(i)
		sb.WriteString("\"")
		if k < len(s) {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}")

	return sb.String()
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for i := range s1 {
		_, ok := s2[i]
		if !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for i := range s1 {
		_, ok := s2[i]
		if ok {
			return false
		}
	}
	for j := range s2 {
		_, ok := s1[j]
		if ok {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	for i := range s1 {
		_, ok := s2[i]
		if !ok {
			return false
		}
	}
	for j := range s2 {
		_, ok := s1[j]
		if !ok {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	s := New()
	for i := range s1 {
		_, ok := s2[i]
		if ok {
			s.Add(i)
		}
	}

	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for i := range s1 {
		_, ok := s2[i]
		if !ok {
			s.Add(i)
		}
	}

	return s
}

func Union(s1, s2 Set) Set {
	s := New()
	for i := range s1 {
		s.Add(i)
	}
	for j := range s2 {
		s.Add(j)
	}
	return s
}
