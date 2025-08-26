package leet

import "sort"

// https://leetcode.com/problems/3sum/description/

// Дан массив nums, состоящий из n целых чисел
// Вернуть все тройки nums[i], nums[j], nums[k], такие что i != j, i != k, j != k и nums[i] + nums[j] + nums[k] == 0
// Решение не должно включать одинаковые тройки

// Идея.
// Массив неупорядочен.
// Можно попробовать упорядочить массив. Чтобы было проще избегать дубликатов.

// Похоже надо создавать три указателя.
// первый фиксируем на нулевом элементе          - p
// второй устанавливаем на первом элементе       - left
// третий в конце                                - right
// первый указатель всегда находится левее остльных

// Если триплет меньше 0, то сдвигаем внутрь массива правый указатель пока значение под указателем right не изменится
// Если триплет больше 0, то сдвигаем внутрь массива левый указатель пока значение под указателем left не изменится
// Если триплет равен 0, то надо сделать следующие действия сдвигаем оба указателя внутрь массива пока значение под указателем left и right не изменится
// Если два указателя встречаются, то надо сделать следующие действия:
// 1. Сдвинуть p указатель внутрь массива пока значение под указателем p не изменится
// 2. Сдвинуть left указатель внутрь массива относительно p
// 3. Установить right указатель на конец массива

//

func threeSum(nums []int) [][]int {

	p := 0
	res := [][]int{}

	sort.Ints(nums)

	for p < len(nums)-2 && nums[p] <= 0 {
		res = append(res, getSum(nums[p+1:], -nums[p])...)
		p = nextLeftPointer(nums, p)
	}

	return res
}

func getSum(arr []int, target int) [][]int {
	left := 0
	right := len(arr) - 1

	res := make([][]int, 0)

	for left < right {
		lval := arr[left]
		rval := arr[right]

		if lval+rval < target {
			left = nextLeftPointer(arr, left)
		} else if lval+rval == target {
			res = append(res, []int{-target, lval, rval})
			left = nextLeftPointer(arr, left)
			right = nextRightPointer(arr, right)
		} else {
			right = nextRightPointer(arr, right)
		}
	}

	return res
}

func nextLeftPointer(arr []int, left int) int {
	prev := arr[left]
	for left < len(arr) && arr[left] == prev {
		left++
	}
	return left
}

func nextRightPointer(arr []int, right int) int {
	prev := arr[right]
	for right > 0 && arr[right] == prev {
		right--
	}
	return right
}
