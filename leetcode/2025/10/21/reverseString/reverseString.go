package reversestring

// https://leetcode.com/problems/reverse-string/?envType=problem-list-v2&envId=two-pointers

// Напишите функцию, которая переворачивает строку. Входная строка задана как массив символов s.
// Это необходимо сделать, изменив входной массив на месте с использованием O(1) дополнительной памяти. Не нужно возвращать исходный массив!
// Example 1:
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]
// Example 2:
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]

// Идея
// используем два указателя first, last
// first указывает на первый элемент
// last указывает на последний элемент
// итерируемся, пока first < last
// меняем местами элементы first и last

func reverseString(s []byte) {
	first := 0
	last := len(s) - 1

	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}
