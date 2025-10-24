package palindromenumber

import "strconv"

// https://leetcode.com/problems/palindrome-number/description/

// Given an integer x, return true if x is a palindrome, and false otherwise.

// Идея
// преобразуем в строку, сравним с перевернутой

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}
