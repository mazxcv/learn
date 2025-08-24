package kindergarten

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// Define the Garden type here.

type Garden struct {
	plants   []string
	children []string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {

	for _, v := range diagram {
		switch v {
		case 'G', 'C', 'R', 'V', '\n':
			continue
		default:
			return &Garden{}, fmt.Errorf("diagram must be 2 string of 'G', 'C', 'R', 'V'")
		}
	}

	plants := strings.Split(diagram, "\n")
	if len(plants) != 3 {
		return &Garden{}, fmt.Errorf("diagram must be 2 string")
	}
	plants = plants[1:]
	if len(plants[0]) != len(plants[1]) {
		return &Garden{}, fmt.Errorf("diagram must be 2 string equals")
	}
	if len(plants[0])%2 != 0 {
		return &Garden{}, fmt.Errorf("diagram must be 2 even number string")
	}
	if len(plants[0])/2 != len(children) {
		return &Garden{}, fmt.Errorf("diagram must be 2 plant for every children")
	}

	uniqueChild := map[string]int{}
	for _, v := range children {
		if _, ok := uniqueChild[v]; ok {
			return &Garden{}, fmt.Errorf("dont unique children")
		}
		uniqueChild[v] = 1
	}

	return &Garden{plants, children}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {

	plant := map[byte]string{
		'G': "grass",
		'C': "clover",
		'R': "radishes",
		'V': "violets",
	}

	children := []string{}
	children = append(children, g.children...)

	sort.Strings(children)

	childIndex := slices.Index(children, child)
	if childIndex == -1 {
		return []string{}, false
	}

	childGarden := []string{
		plant[g.plants[0][childIndex*2]],
		plant[g.plants[0][childIndex*2+1]],
		plant[g.plants[1][childIndex*2]],
		plant[g.plants[1][childIndex*2+1]],
	}

	return childGarden, true
}
