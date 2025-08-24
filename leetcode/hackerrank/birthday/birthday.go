package birthday

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d - sum of values in s
 *  3. INTEGER m - len sub
 */

func birthday(s []int32, d int32, m int32) (res int32) {
	if m > int32(len(s)) {
		return 0
	}

	var i int32
	var sum int32
	for i = 0; i < m; i++ {
		sum += s[i]
	}

	if sum == d {
		res++
	}

	for i = m; i < int32(len(s)); i++ {
		sum = sum + s[i] - s[int(i)-int(m)]
		if sum == d {
			res++
		}
	}
	return
}
