package example

// https://leetcode.com/problems/move-zeroes/description/

// Дан массив nums, переместить все нули в конец, сохраняя порядок остальных элементов
// Надо делать in-place без использования дополнительного пространства

// Будем использовать подход с двумя указателями
// right - pointer для поиска ненуделвый элементов
// left - pointer для отслеживания, куда надо поместить следующий ненулевой элемент
// Когда right находит ненулевой элемент, мы меняем элементы местами с left
// Перемещаем оба указателя  вперед
// Как только right достигнет конца, мы останавливаемся и возвращаем массив

func moveZeroes(nums []int) []int {
	left := 0
	for right := range len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	return nums
}
