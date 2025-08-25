package leet

// https://leetcode.com/problems/reverse-bits/description/

// Дано целое число n, вернуть число с битами в обратном порядке

// Идея.
// Пробовать сдвигать n вправо и получленный бит добавлять в начало результата
// Делать так столько раз, какова разрядность типа int

func reverseBits(n int) int {
	res := 0

	for range 32 {
		bit := n & 1
		res = (res << 1) | bit
		n >>= 1
	}

	return res
}
