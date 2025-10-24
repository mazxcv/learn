package searchinsertposition

// https://leetcode.com/problems/search-insert-position/description/

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.
// Example 1:
// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:
// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:
// Input: nums = [1,3,5,6], target = 7
// Output: 4

// Идея
// проверим границы
// если target меньше nums[0], то возвращаем 0
// если target больше nums[len(nums)-1], то возвращаем len(nums)
// В остальных случаях методом половинного деления надо выяснить позицию вставки
// Нужны 2 индекса l и r

// Делим пополам
// Если target < nums[m], то ищем в левой половине
// Если target > nums[m], то ищем в правой половине
// Если target == nums[m], то двигаемся влево, пока не найдем элемент меньше target
// возвращаем индекс позиции вставки + 1
// Если l == r, то возвращаем l

func searchInsert(nums []int, target int) int {
	index := 0
	start_ind := 0
	end_ind := len(nums) - 1

	for start_ind <= end_ind {
		mid_ind := start_ind + (end_ind-start_ind)/2
		if nums[mid_ind] == target {
			return mid_ind
		} else if nums[mid_ind] < target {
			start_ind = mid_ind + 1
			index = start_ind
		} else {
			end_ind = mid_ind - 1
			index = mid_ind
		}
	}
	return index
}
