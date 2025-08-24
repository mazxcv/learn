package prime

import "fmt"

func MaybePrime(lMaybyPrime int, aPrime []int) bool {
	for _, v := range aPrime {
		if lMaybyPrime%v == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (prime int, err error) {
	if n <= 0 {
		return 0, fmt.Errorf("n must be pozitive non zero value")
	}

	switch n {
	case 1:
		prime = 2
	default:
		aPrimeList := []int{2}
		lMaybyPrime := 3
		for i := 2; i <= n; i++ {
			for !MaybePrime(lMaybyPrime, aPrimeList) {
				lMaybyPrime++
			}
			aPrimeList = append(aPrimeList, lMaybyPrime)
		}
		prime = lMaybyPrime
	}
	return
}
