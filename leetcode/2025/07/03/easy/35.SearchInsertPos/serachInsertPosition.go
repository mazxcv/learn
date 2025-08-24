package searchinsertpos

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}
	if target <= nums[0] {
		return 0
	}

	l := 0
	r := len(nums) - 1
	m := (l + r) / 2

	for {
		switch {
		case target == nums[m]:
			return m
		case target == nums[r]:
			return r
		case target == nums[l]:
			return l

		case r == l:
			if target > nums[r] {
				return r + 1
			}
			return l
		case r == l+1:
			if target > nums[r] {
				return r + 1
			}
			if target > nums[l] {
				return r
			}
			return l
		case r == l+2:
			if target > nums[r] {
				return r + 1
			}
			if target > nums[m] {
				return r
			}
			if target > nums[l] {
				return m
			}
			if target < nums[l] {
				return l
			}
		case target > nums[m]:
			l = m
			r = r - 1
			m = (l + r) / 2
		case target < nums[m]:
			l = l + 1
			r = m
			m = (l + r) / 2
		}
	}
}
