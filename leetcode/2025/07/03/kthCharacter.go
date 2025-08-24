package kthcharacter

import "math"

// https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/?envType=daily-question&envId=2025-07-03
// Задача найти к-ый символ в строке, которая генерируется по следующему условию:
// Исходные данные: строка с символом 'a'
// Количество итераций генерации - k
// Используется английский алфавит в нижнем регистре ('a' 'b' 'c' 'd' 'e' 'f' 'g' 'h' 'i' 'j' 'k' 'l' 'm' 'n' 'o' 'p' 'q' 'r' 's' 't' 'u' 'v' 'w' 'x' 'y' 'z')
// Множесто является кольцом, тоесть после 'z' -> 'a'
//
// Каждая генерация берет строку из предыдущей генерации, для каждого символа берет следующий символ в алфавите, если он больше 'z', то берет 'a'.
// Полученная строка добавляется к к исходной  строке.
// Требуется вернуть k-й символ строки после k генераций

// Пример:
// k = 5
// a -> ab -> abbc -> abbcbccd
// k-й по счету в строке - это k-1 индекс, поэтому Результат: 'b'

// 1. Идея:
// Исходя из удвоений длины строки после каждой генерации, можно сделать log(k) генераций, чтобы получить нужный символ

func kthCharacter(k int) byte {
	if k <= 1 {
		return 'a'
	}

	currentString := "a"
	power := int(math.Log2(float64(k))) + 2
	for i := 1; i < power; i++ {
		currentString = generateString(currentString)
	}

	return currentString[k-1]
}

// generateString takes a string and returns a new string with double the length,
// where each character is replaced by itself followed by the next character in
// the English alphabet sequence. If the last character is 'z', it wraps around and
// becomes 'a'.
func generateString(input string) string {
	result := make([]byte, len(input)*2)
	for index, char := range input {
		result[index] = byte(char)
		result[index+len(input)] = nextChar(byte(char))
	}

	return string(result)
}

// nextChar returns the next character in the English alphabet sequence.
// If the current character is 'z', it wraps around and returns 'a'.
func nextChar(currentChar byte) byte {
	if currentChar == 'z' {
		return 'a'
	}
	return currentChar + 1
}
