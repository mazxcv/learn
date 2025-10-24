package checkifdigitsareequalinstringafteroperationsi

// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/description/

// Дана строка из цифр
// Для каждой пары цифр необходимо, начиная с первой, сложить по модулю 10
// С получившейся строкой повторить действия пока она не превратится в 2 цифры
// Если цифры равны вернуть true, иначе false

// Идея:
// Нет здесь идеи - надо делать поэтапно и каждый раз вычислять новую строку

func hasSameDigits(s string) bool {
	if len(s) == 1 {
		return true
	}

	ints := make([]int, len(s))

	for i := range len(s) {
		ints[i] = int(s[i] - '0')
	}

	for len(ints) != 2 {
		for i := range len(ints) - 1 {
			ints[i] = (ints[i] + ints[i+1]) % 10
		}
		ints = ints[:len(ints)-1]
	}

	return ints[0] == ints[1]
}
