package example

// https://leetcode.com/problems/range-sum-query-immutable/description/

// Дан целочисленный массив nums

// 1. Вычислить сумму элементов между индексами left и right где 0 <= left <= right < len(nums)
// 2. Реализовать class NumArray с конструктором Constructor(int[] nums)
// и методом int sumRange(int left, int right), который возвращает сумму элементов между индексами left и right включительно
//

// Преобразуем входной массив в массив префиксных сумм
// Создадим структуру NumArray который включает в себя массив префиксной суммы
// Когда потребуется вычислить сумму элементов между индексами left и right, то возвращаем nums[right] - nums[left]

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	na := &NumArray{
		nums: nums,
	}

	return *na
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}
	return this.nums[right] - this.nums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
