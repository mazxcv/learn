package prime

func Factors(n int64) (primes []int64) {
	if n <= 1 {
		return []int64{}
	}

	var i int64 = 2
	for n != 1 {
		for n%i == 0 {
			primes = append(primes, i)
			n /= i
		}
		i++
	}
	return
}
