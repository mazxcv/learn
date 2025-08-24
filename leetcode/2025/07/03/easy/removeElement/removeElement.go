package removeelement

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	lastIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[lastIndex] = nums[i]
			lastIndex++
		}
	}

	return lastIndex
}
