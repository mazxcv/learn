package leet

// https://leetcode.com/problems/number-of-1-bits/description/

// Дано целое число n, вернуть количество единиц в его двоичном представлении
// Вес Хэмминга - это количество единиц в двоичном представлении числа

// Идея.
// пробовать сдвигать n вправо и считать количество единиц

func hammingWeight(n int) int {
	res := 0
	for n > 0 {
		res += n & 1
		n >>= 1
	}

	return res
}
