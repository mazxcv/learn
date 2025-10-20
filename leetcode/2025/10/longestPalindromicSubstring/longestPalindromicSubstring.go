package longestpalindromicsubstring

// https://leetcode.com/problems/longest-palindromic-substring/description/?envType=problem-list-v2&envId=two-pointers

// Найти самыую длинную подстроку, которая является палиндромом в заданной строке
// Example 1:
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.

// Example 2:
// Input: s = "cbbd"
// Output: "bb"

// Идея
// Начиная с начала ищем максимально возможный палиндром, запоминаем
// двигаемся вперед на один символ, опять ищем палиндром, если получили большийпо размеру, то запоминаем

// Как искать палиндром?
// 1. нечетный палиндром
// 1.1. отступаем вправо на один символ, смотрим их равенство начала и следующего
// 1.2 отступаем влево и вправо на один символ, смотрим их равенство
// 1.3. Делаем так, пока получется палиндром

// 2. четный палиндром
// 2.1. отступаем влево и вправо на один символ, смотрим их равенство
// 2.2. Делаем так, пока получется палиндром

// Видно, что алгоритмы четного и нечетного палиндрома почти совпадают. Поэтому можно объединить
// Если следующий символ не равен текущему, то пробуем проверять нечетный палиндром

func longestPalindrome(s string) string {
	// Двигаемся центром вдоль строки и  запоминаем наиболее достойного кандидата
	candidate := ""
	for cntr := range len(s) {
		newCandidate := getCandidatePalindrome(cntr, s)
		if len(newCandidate) > len(candidate) {
			candidate = newCandidate
		}
	}

	return candidate
}

// Для каждого центра ищем границы возможного палиндрома
func getBoundaryOfPalindrome(cntr int, left int, right int, s string) (int, int) {
	for {
		if left == 0 || right == len(s)-1 {
			break
		}
		if s[left-1] != s[right+1] {
			break
		}

		left--
		right++
	}

	return left, right
}

func getCandidatePalindrome(cntr int, s string) string {

	left, right := getBoundaryOfPalindrome(cntr, cntr, cntr, s)

	if cntr+1 < len(s) && s[cntr+1] == s[cntr] {
		leftCandidate, rightCandidate := getBoundaryOfPalindrome(cntr, cntr, cntr+1, s)
		if rightCandidate-leftCandidate > right-left {
			left = leftCandidate
			right = rightCandidate
		}
	}

	return s[left : right+1]
}
