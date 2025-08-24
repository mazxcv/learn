package findlucky

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	resutlt := -1
	for k, v := range m {
		if k == v && resutlt < k {
			resutlt = k
		}
	}
	return resutlt
}
