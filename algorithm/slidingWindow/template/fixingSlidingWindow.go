package template

func fixingSlidingWindow(arr []int, k int) int {
	windowSum := 0
	var requiredValue int

	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	requiredValue = windowSum

	// Slide the window across the array
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		requiredValue = compute(requiredValue, windowSum, predicate)
	}

	return requiredValue
}

func predicate(windowSum int, requiredValue int) bool {
	return windowSum > requiredValue
}

func compute(requiredValue int, windowSum int, f func(int, int) bool) int {
	if f(windowSum, requiredValue) {
		return windowSum
	}
	return requiredValue
}
