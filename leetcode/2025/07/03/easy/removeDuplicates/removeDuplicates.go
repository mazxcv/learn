package removeduplicates

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	lastIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lastIndex] {
			lastIndex++
			nums[lastIndex] = nums[i]
		}
	}

	return lastIndex + 1
}
