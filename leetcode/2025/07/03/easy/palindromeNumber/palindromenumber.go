package palindromenumber

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	var y int
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	return x == y || x == y/10

}

func IsPalindromeNaive(x int) bool {

	if x < 0 {
		return false
	}

	remainder := []int{}
	for x > 0 {
		remainder = append(remainder, x%10)
		x = x / 10
	}

	for i := 0; i < len(remainder)/2; i++ {
		if remainder[i] != remainder[len(remainder)-i-1] {
			return false
		}
	}
	return true

}
