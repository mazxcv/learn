// package proberb is a small library for generate  proverb
package proverb

import "fmt"

// func Proverb generate the relevant proverb
func Proverb(rhyme []string) (layout []string) {
	if len(rhyme) < 1 {
		return
	}
	if len(rhyme) == 1 {
		layout = append(layout, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		return
	}
	for i := 0; i < len(rhyme)-1; i++ {
		layout = append(layout, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	layout = append(layout, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return

}
