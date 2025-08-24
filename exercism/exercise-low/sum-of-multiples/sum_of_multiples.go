package summultiples

func SumMultiples(limit int, divisors ...int) (energy int) {

	if limit <= 1 {
		return 0
	}
	for i := 1; i < limit; i++ {
		for _, v := range divisors {
			if v == 0 {
				continue
			}
			if i%v == 0 {
				energy += i
				break
			}
		}
	}

	return

}
