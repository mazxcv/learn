package divisiblesumpairs

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n - len(ar)
 *  2. INTEGER k - на  сколько нужно делить число
 *  3. INTEGER_ARRAY ar
 * основная идея - собрать мап с остатками от делений на k, затем найти пары внутри мапы
 * и обнаружить сколько раз можно пару собрать  - произведение  встречаемого,
 * в случае равных - надо найти сумму всех попарных соединений, тоесть сумму от 1 до k-1
 */

func divisibleSumPairs(n int32, k int32, ar []int32) (res int32) {
	if k == 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 0
	}

	if k < 0 {
		k = -k
	}

	remainder := map[int32]int32{}
	var i int32

	for i = 0; i < n; i++ {
		val := ar[i] % k
		if val < 0 {
			val += k
		}
		remainder[val]++
	}

	for i = 0; i <= k/2; i++ {
		switch {
		case i == 0:
			res += remainder[i] * (remainder[i] - 1) / 2
		case i == k-i:
			res += remainder[k/2] * (remainder[k/2] - 1) / 2
		default:
			res += remainder[i] * remainder[k-i]
		}
	}
	return
}
