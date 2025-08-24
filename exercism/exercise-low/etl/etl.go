package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}
	for k, v := range in {
		for _, s := range v {
			res[strings.ToLower(s)] = k
		}
	}
	return res
}
