package leet

// https://leetcode.com/problems/permutation-in-string/description/

// Дано две строки s1 и s2, вернуть true, если s2 содержит перестановку s1
// Перестановка - это слово, состоящее из одних и тех же букв, но в разном порядке
// s1, s2 - строки символов латинского алфавита в нижнем регистре

// Идея.
// Используем скользящее окно размером s1.
// В виду того, что сипользуются символы латинского алфавита в нижнем регистре, можно использовать массив частот символов.
// Строим первое окно.
// Строим массив частот символов в окне.
// Строим массив частов символов в s1.

// Реализуем метод сравнения массивов частот символов для s1 и окна. O(n)

// реализуем безопасное удаление символа из массива частот символов. O(1)
// реализуем добавление символа в массив частот символов. O(1)

const MAX_CHAR = 26

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	k := len(s1)
	freqS1 := make([]int, MAX_CHAR)
	for i := range k {
		safeAddChar(freqS1, s1[i])
	}

	freqS2 := make([]int, MAX_CHAR)
	for i := range k {
		safeAddChar(freqS2, s2[i])
	}

	if isPermutation(freqS1, freqS2) {
		return true
	}

	for i := k; i < len(s2); i++ {
		safeRemoveChar(freqS2, s2[i-k])
		safeAddChar(freqS2, s2[i])

		if isPermutation(freqS1, freqS2) {
			return true
		}
	}
	return false

}

func isPermutation(a, b []int) bool {
	for i := range MAX_CHAR {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func safeRemoveChar(a []int, c byte) {
	if a[c-'a'] > 0 {
		a[c-'a']--
	}
}

func safeAddChar(a []int, c byte) {
	a[c-'a']++
}
