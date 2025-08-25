package example

// https://leetcode.com/problems/container-with-most-water/

// Данн массив целых чисел height, состоящий из n элементов.
// На нем нарисованы n вертикальных линий, высотой соответствующего элемента массива.
// Найти две линии, которые могут удерживать максимальной количество воды между собой
// Тоесть найти контейнер с максимальной площадью воды

// S зависит от двух факторов - высота (минимальная) и расстояние между линиями
// Площадь S = высота * расстояние между линиями

// Идея. Не проверять каждую пару
// Начать с максимально широкого контейнера
// Постепенно уменьшать ширину, перемещая указатели внутрь, чтобы вычислить можно ли сформровать большую площадь

// left размещается в начале
// right размещается в конце
// MaxArea = 0
// Движемся пока не встрется right left
// Рассчитываем площадь
// Перемещаем указатель на меньшую высоту внутрь

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1
	maxArea := 0

	for left <= right {
		width := right - left
		minHeight := min(height[left], height[right])
		area := minHeight * width
		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
