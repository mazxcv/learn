package example

// https://leetcode.com/problems/maximum-average-subarray-i/description/
// Дан целочисленный массив nums, состоящий из n элементов, и целое число k.
// Найти непрерывный подмассив, длина которого равна k, который имеет максимальное среднее значение и вернуть его

// Идея.
// Использовать скользящее окно размером k, которое передвигается по одному элементу
// каждый раз обновляет сумму элементов в окне.
// Находим максимум.
// Считаем среднее.

func findMaxAverage(nums []int, k int) float64 {

	windowState := 0

	if k > len(nums) {
		return 0
	}

	for i := range k {
		windowState += nums[i]
	}
	res := windowState

	for i := k; i < len(nums); i++ {
		windowState += nums[i] - nums[i-k]
		res = getResult(res, windowState)
	}

	return float64(res) / float64(k)
}

func getResult(prev, current int) int {
	if current > prev {
		return current
	}
	return prev
}
