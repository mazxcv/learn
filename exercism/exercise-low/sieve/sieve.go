package sieve

// type SieveEratosphen
type SieveEratosphen struct {
	Struct map[int]bool
	limit  int
}

// func NewSieveEratosphen - prepare new Eratosphen
func NewSieveEratosphen(limit int) SieveEratosphen {
	if limit <= 1 {
		return SieveEratosphen{}
	}
	var preparePrime = make(map[int]bool, limit-2)

	for i := 2; i <= limit; i++ {
		preparePrime[i] = true
	}
	return SieveEratosphen{preparePrime, limit}
}

// func Seed - Mark all the multiples of that key number as not prime
func (e SieveEratosphen) Seed(key int) {
	for i := key * 2; i <= e.limit; i += key {
		e.Struct[i] = false
	}
}

// func Prepare - marked all key isPrime
func (e SieveEratosphen) Prepare() SieveEratosphen {
	for k := 2; k <= e.limit; k++ {
		if e.Struct[k] {
			e.Seed(k)
		}
	}
	return e
}

// func GetPrimeList - prepare list of Prime
func (e SieveEratosphen) GetPrimeList() []int {
	t := []int{}
	for k := 2; k <= e.limit; k++ {
		if e.Struct[k] {
			t = append(t, k)
		}
	}
	return t
}

func Sieve(limit int) []int {
	return NewSieveEratosphen(limit).Prepare().GetPrimeList()
}
