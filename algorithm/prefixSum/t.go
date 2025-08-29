package prefixsum

func calcPrefixSum(num []int) []int {
	pSum := make([]int, len(num))

	pSum[0] = num[0]

	for i := range num[1:] {
		pSum[i] = pSum[i-1] + num[i]
	}

	return pSum
}
