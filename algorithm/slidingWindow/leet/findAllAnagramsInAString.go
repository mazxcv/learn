package leet

// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

// Дано 2 строки s и p
// Вернуть список индексов, в которых начинаются анаграммы p в s
// Анаграмма - это слово, состоящее из одних и тех же букв, но в разном порядке

// s, p - строки символов латинского алфавита в нижнем регистре

// Идея.
// Используем скользящее окно размером с длину p.
// Построим частотный словарь символов в p.
// Строим первое окно.
// Строим частотный словарь символов в окне.

// реализуем метод стравнения частотных словарей для p и окна. O(n)
// Если совпадают - добавляем индекс в результат

func findAnagrams(s string, p string) []int {

	k := len(p)

	if len(p) > len(s) {
		return nil
	}

	freqP := make(map[byte]int)
	for i := range k {
		addChar(freqP, p[i])
	}

	freqS := make(map[byte]int)
	for i := 0; i < k; i++ {
		freqS[s[i]]++
	}

	var res []int

	if isAnagram(freqP, freqS) {
		res = append(res, 0)
	}

	for i := k; i < len(s); i++ {
		removeChar(freqS, s[i-k])
		addChar(freqS, s[i])

		if isAnagram(freqP, freqS) {
			res = append(res, i-k+1)
		}
	}
	return res
}

func isAnagram(s, p map[byte]int) bool {
	if len(s) != len(p) {
		return false
	}

	for k, v := range p {
		if s[k] != v {
			return false
		}
	}

	return true
}

func removeChar(s map[byte]int, c byte) {
	if v, ok := s[c]; ok {
		if v == 1 {
			delete(s, c)
		} else {
			s[c]--
		}
	}
}

func addChar(s map[byte]int, c byte) {
	if _, ok := s[c]; ok {
		s[c]++
	} else {
		s[c] = 1
	}
}
