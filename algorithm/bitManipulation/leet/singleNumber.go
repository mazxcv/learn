package leet

// https://leetcode.com/problems/single-number/description/

// Дан непустой массив целых чисел nums, где все элементы встречаются дважды, кроме одного. Найти этот один элемент.

// Идея.
// Использовать битовый оператор xor ^ Сложность O(n)
// если элементы встречаются дважды, то xor будет равен 0
// Поскольку только один элемент встречается один раз, то xor будет равен этому элементу

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}
