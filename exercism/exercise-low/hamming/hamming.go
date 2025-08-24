// Package hamming is a small library for determining the hamming distance between two strands.
package hamming

import "fmt"

// Distance calculates Hamming distance for 2 DNA strands
// accepts two strands and returns their hamming distance.
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		err = fmt.Errorf("length strands not equal")
		return
	}
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return
}
