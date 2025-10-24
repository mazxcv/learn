package iintersectionoftwoarrays

import "sort"

// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.
// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]
// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.

// Идея
// сортировка + указатель
// тогда сложность O(nlog(n))
// Правила движения:
// сравниваем значения под указателями
// если равны, то пушим элемент в результирующий массив и двигаем указатели, пока не изменятся оба
// если не равны, то двигаем указатель, который меньше
// и при этом второй указатель станет указывать на большее число или тоже дойдет до конца

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)

	sort.Ints(nums1)
	sort.Ints(nums2)

	p1 := 0
	p2 := 0

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			intersect := nums1[p1]
			res = append(res, intersect)
			for p1 < len(nums1) && nums1[p1] == intersect {
				p1++
			}
			for p2 < len(nums2) && nums2[p2] == intersect {
				p2++
			}
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		}
	}

	return res
}
