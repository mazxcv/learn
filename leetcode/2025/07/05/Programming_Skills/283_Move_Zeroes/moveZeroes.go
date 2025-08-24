package _83_Move_Zeroes

import "fmt"

func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			if nums[l] != nums[r] {
				nums[l] = nums[r]
				nums[r] = 0

			}
			l++
		}
		r++
	}
	fmt.Println(nums)
}
