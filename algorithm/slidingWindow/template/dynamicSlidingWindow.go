package template

func dynamicSlidingWindow(arr []int, k int) int {
	windowState := 0
	result := 0

	for right := 0; right < len(arr); right++ {
		windowState = calculateWindow(windowState, arr[right])

		if !predicateD(arr[right]) {
			result = res(result, windowState)
			windowState = 0
		}
	}

	return result
}

func calculateWindow(prevState, currentValue int) int {
	return prevState + currentValue
}

func predicateD(currentValue int) bool {
	return currentValue > 0
}

func res(prevResult, currentValue int) int {
	if currentValue > prevResult {
		return currentValue
	}
	return prevResult
}
