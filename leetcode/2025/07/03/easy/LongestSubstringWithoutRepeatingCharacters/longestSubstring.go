package longestsubstringwithoutrepeatingcharacters

import "strings"

// Идея состоит в том, что надо пройтись по строке,  пряча все руны в map в качестве ключа, значение индекс строки O(n)
// когда встречаем уже существующую руну, разделяем поток исполнения
// если разность индексов + 1 сравнима с текущей длиной, значит длина подстроки имеет шанс увеличиться, а значит надо запомнить следующий после дублированного индекса, заменить в map на новый индекс, продолжить перебор рун
// если разность индексов + 1 меньше текущей длины, значит у подстроки нет шансов увеличиться по предыдущему подходу, значит надо удалить все руны из map до найденного индекса
// если длина строки минус индекс текущей руны меньше длины уже известной максимальной подстроки, значит мы нашли длину максимальной подстроки
func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	first := 0
	maxLen := 0
	for i, v := range strings.Split(s, "") {
		if index, ok := m[v]; ok {
			if index == first {
				m[v] = i
				first++
			} else {
				if maxLen < len(m) {
					maxLen = len(m)
				}
				substr := strings.Split(s[first:index], "")
				for _, j := range substr {
					delete(m, j)
				}
				m[v] = i
				first = index + 1
			}
		}
		m[v] = i
	}

	if maxLen < len(m) {
		maxLen = len(m)
	}

	return maxLen
}
