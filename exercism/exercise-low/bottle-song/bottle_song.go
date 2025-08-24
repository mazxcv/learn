package bottlesong

import (
	"fmt"
	"strings"
)

func IntToWord(dec int) string {
	return map[int]string{
		10: "Ten green bottles",
		9:  "Nine green bottles",
		8:  "Eight green bottles",
		7:  "Seven green bottles",
		6:  "Six green bottles",
		5:  "Five green bottles",
		4:  "Four green bottles",
		3:  "Three green bottles",
		2:  "Two green bottles",
		1:  "One green bottle",
		0:  "No green bottles",
	}[dec]
}

// Ten green bottles hanging on the wall,
// Ten green bottles hanging on the wall,
// And if one green bottle should accidentally fall,
// There'll be nine green bottles hanging on the wall.
//
// Nine green bottles hanging on the wall,
// Nine green bottles hanging on the wall,
// And if one green bottle should accidentally fall,
// There'll be eight green bottles hanging on the wall.
func Recite(startBottles, takeDown int) (song []string) {
	if takeDown > startBottles {
		return []string{}
	}

	i := startBottles
	for i > startBottles-takeDown {
		song = append(song, fmt.Sprintf("%s hanging on the wall,", IntToWord(i)))
		song = append(song, fmt.Sprintf("%s hanging on the wall,", IntToWord(i)))
		song = append(song, "And if one green bottle should accidentally fall,")
		song = append(song, fmt.Sprintf("There'll be %s hanging on the wall.", strings.ToLower(IntToWord(i-1))))
		i--
		if i != startBottles-takeDown {
			song = append(song, "")
		}
	}
	return
}
