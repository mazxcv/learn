package example

// https://leetcode.com/problems/subarray-sum-equals-k/description/

// Дан целочисленный массив nums и целое число k.
// Найти общее количество подмассивов, сумма которых равна k

// Все подмассивы - не пусты

// Идея.
// Подготовим массив префиксых сумм массива
// сумма каждого подмассива равна сумме элементов между индексами left и right
// Для массива префиксных сумм это означает разность префиксных сумм между индексами
// left(надо брать left-1, если это невозможно, то значение ноль) и right
// rPrefix - lPrefix = k -> rPrefix -k = lPrefix
// Тогда задача сведется к тому, чтобы посчитать сколько раз встретилась rPrefix -k
// Для хранения частоты каждого префикса надо использовать частотную мапу

func subarraySum(nums []int, k int) int {

	prefixSum := 0
	count := 0

	prefixSumFreq := make(map[int]int)
	prefixSumFreq[0] = 1

	for _, num := range nums {
		prefixSum += num
		count += prefixSumFreq[prefixSum-k]

		prefixSumFreq[prefixSum]++
	}
	return count
}
