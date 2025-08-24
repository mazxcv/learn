package tupleSameProduct

func tupleSameProduct(nums []int) int {

	result := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			result[nums[i]*nums[j]]++
		}
	}

	total := 0
	for _, v := range result {
		if v > 1 {
			total += 8 * (v * (v - 1) / 2)
		}
	}

	return total
}
