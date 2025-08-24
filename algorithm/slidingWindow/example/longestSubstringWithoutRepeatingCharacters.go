package example

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

// Дана строка s, вернуть максимальную длину подстроки, не содержащую повторяющиеся символы
// Тоесть подстрока содержит только уникальные символы

// Идея.
// Строим первое скользящее окно, пока не встретим повторяющийся символ.
// Используем два указателя left и right.
// Сначала двигаем right, пока не встретим повторяющийся символ.
// Для эффективной проверки на дубликат будет использовать частотный словарь.

// Сначала оба указателя left и right будут указывать нулевой индекс
// Добавим в словарь упоминание первого символа
// Пока не  закончилось строка:
// двигаем right
// Добавляем символы в словарь
// Пока нарушается условие уникальности:
// Обновить максимальную длину
// Убираем символы под индексом left из словаря
// Двигаем left.
//

func lengthOfLongestSubstring(s string) int {
	windowState := make(map[byte]bool)
	res := 0
	windowLen := 0
	left := 0

	for right := range len(s) {
		for windowState[s[right]] {
			delete(windowState, s[left])
			left++
			windowLen--
		}
		windowState[s[right]] = true
		windowLen++
		res = max(res, windowLen)
	}

	return res

}

func max(prev, current int) int {
	if current > prev {
		return current
	}
	return prev
}
