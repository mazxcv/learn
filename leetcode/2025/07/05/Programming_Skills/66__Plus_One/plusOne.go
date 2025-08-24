package _6__Plus_One

func plusOne(digits []int) []int {
	p := 1
	for i := len(digits) - 1; i >= 0; i-- {
		res := digits[i] + p
		digits[i] = res % 10
		p = res / 10
	}
	if p != 0 {
		arr := []int{p}
		arr = append(arr, digits...)
		return arr
	}
	return digits
}
