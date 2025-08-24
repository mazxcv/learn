package armstrong

func SumofCube(in []int) (t int) {
	for _, v := range in {

		temp := 1
		for i := 1; i <= len(in); i++ {
			temp *= v
		}

		t += temp
	}
	return
}

func IsNumber(n int) bool {

	if n <= 0 {
		return true
	}
	t := n
	digits := []int{}
	for t > 0 {
		digits = append(digits, t%10)
		t /= 10
	}
	return n == SumofCube(digits)
}
