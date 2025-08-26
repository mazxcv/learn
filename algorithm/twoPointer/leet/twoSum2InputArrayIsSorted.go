package leet

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

// Дан массив отсортированный по неубыванию.
// Вернуть какие по счету элементы в сумме дают target,
// Тоесть  i и j, такие что numbers[i] + numbers[j] = target  и вернуть [i+1, j+1]
// Нельзя использовать элемент дважды.

// Идея.
// Использовать два указателя. Первый находится на нулевой позиции
// Второй на последней
// Передвигаем указатели по следующий условиям:
// Если сумма элементов больше target, то сдвигаем внутрь массива правый указатель
// Если сумма элементов меньше target, то сдвигаем внутрь массива левый указатель
// Если сумма элементов равна target, то возвращаем индексы, увеличенные на единицу,
// потому что есть nребование вернуть какие по счету элементы дают нужную сумму

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}
