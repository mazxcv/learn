package breakingbestandworstrecord

func breakingRecords(scores []int32) []int32 {

	if len(scores) == 0 {
		return []int32{0, 0}
	}

	var max, min int32 = scores[0], scores[0]
	var cmax, cmin int32

	for _, v := range scores {
		if max < v {
			cmax++
			max = v
		}
		if min > v {
			cmin++
			min = v
		}
	}
	return []int32{cmax, cmin}

}
