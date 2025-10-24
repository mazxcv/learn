package longestpalindrome

// https://leetcode.com/problems/longest-palindrome/description/

// Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
// Letters are case sensitive, for example, "Aa" is not considered a palindrome.
// Example 1:
// Input: s = "abccccdd"
// Output: 7
// Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
// Example 2:
// Input: s = "a"
// Output: 1
// Explanation: The longest palindrome that can be built is "a", whose length is 1.

// Constraints:
// 1 <= s.length <= 2000
// s consists of lowercase and/or uppercase English letters only.

// Идея
// палиндром - читается одинаково слева направо и справа налево
// надо разобрать все буквы по парам
// плюс можно добавить 1 букву непарную в центр, если таковая существует
// Алгоритм:
// хранилище букв с их количеством - db
// число букв с нечетным количеством - odd_cnt
// результат - res
// Цикл по массиву
// Добавляем букву в хранилище
// если количество букв четное, то res += 2, odd_cnt -= 1
// если количество букв нечетное, то odd_cnt += 1
// в конце цикла проверяем odd_cnt, если оно > 0, то res++

func longestPalindrome(s string) int {
	res := 0
	odd_cnt := 0
	db := make(map[rune]int)

	for _, v := range s {
		db[v]++
		if db[v]%2 == 0 {
			res += 2
			odd_cnt--
		} else {
			odd_cnt++
		}
	}

	if odd_cnt > 0 {
		res++
	}

	return res
}
