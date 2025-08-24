package _36__Single_Number

func singleNumber(nums []int) (res int) {
	for _, v := range nums {
		res ^= v
	}
	return
}
