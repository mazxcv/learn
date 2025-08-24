package migratorybirds

func migratoryBirds(arr []int32) (minType int32) {
	birdsMap := map[int32]int32{}

	for _, v := range arr {
		birdsMap[v]++
	}

	var maxVal int32
	for k, v := range birdsMap {

		switch {
		case v > maxVal:
			minType = k
			maxVal = v
		case v == maxVal:
			if minType > k {
				minType = k
			}
		}
	}
	return
}
