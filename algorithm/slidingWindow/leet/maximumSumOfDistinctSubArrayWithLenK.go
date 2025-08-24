package leet

// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/

// Дан целочисленный массив nums, состоящий из n элементов, и целое число k.
// Найти максимальную сумму непрерывного подмассива с длиной k, все элементы которого уникальны.
// Если таких подмассивов нет, вернуть 0.

// Идея.
// Использовать скользящее окно c двумя условиями: размер окна k, уникальность элементов в окне.
// Введем словарь индексации чисел для окна, в качестве ключа хранится число, а в качестве значения его индекс в окне.
// Цикл по всем элементам массива.
// Если условие нарушилось для правой границы, значит в словаре индексации по числу хранится индекс будущего левого края окна
// значит из суммы по окну вычитаем все элементы от текущего левого окна до индекса дубликата в словаре
// Такжу из словаря индексов удаляем встречаемые числа

func maximumSubarraySum(nums []int, k int) int64 {
	if k > len(nums) {
		return 0
	}

	windowSum := int64(0)
	left := 0
	res := int64(0)
	freq := make(map[int]int)

	for right := range nums {
		if v, ok := freq[nums[right]]; ok {
			for left <= v {
				windowSum -= int64(nums[left])
				delete(freq, nums[left])
				left++
			}
		}
		windowSum += int64(nums[right])
		freq[nums[right]] = right

		if right-left+1 == k {
			res = max(res, windowSum)
			// сдвигаем левую границу
			windowSum -= int64(nums[left])
			delete(freq, nums[left])
			left++
		}

	}
	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
