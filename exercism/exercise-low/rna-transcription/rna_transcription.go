package strand

import "strings"

// func ToRNA determine the RNA complement of a given DNA sequence
func ToRNA(dna string) string {
	dnaToRNA := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	var sb strings.Builder
	sb.WriteString("")
	for _, v := range dna {
		sb.WriteRune(dnaToRNA[v])
	}

	return sb.String()
}
