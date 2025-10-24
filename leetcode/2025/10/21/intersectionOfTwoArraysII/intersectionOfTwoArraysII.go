package intersectionoftwoarraysii

import "sort"

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.

// Constraints:
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000

// Идея
// Сортировка + указатель
// тогда сложность O(nlog(n))
// Правила движения:
// сравниваем значения под указателями
// если равны, то пушим элемент в результирующий массив
// если не равны, то двигаем указатель, который меньше
// и при этом второй указатель станет указывать на большее число или тоже дойдет до конца

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	p1, p2 := 0, 0
	res := make([]int, 0)

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			res = append(res, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		}
	}

	return res
}
