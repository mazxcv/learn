package _96__Monotonic_Array

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	var mono int
	switch {
	case nums[0] < nums[len(nums)-1]:
		mono = 1
	case nums[0] > nums[len(nums)-1]:
		mono = -1
	default:
		mono = 0
	}

	for i := 1; i < len(nums); i++ {
		switch mono {
		case 0:
			if nums[i] != nums[i-1] {
				return false
			}
		case 1:
			if nums[i] < nums[i-1] {
				return false
			}
		default:
			if nums[i] > nums[i-1] {
				return false
			}
		}
	}
	return true
}
