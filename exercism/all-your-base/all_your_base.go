package allyourbase

import (
	"fmt"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	if inputBase < 2 {
		return []int{}, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, fmt.Errorf("output base must be >= 2")
	}

	sum := 0
	for i, v := range inputDigits {
		if v < 0 || v >= inputBase {
			return []int{}, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		sum += v * int(math.Pow(float64(inputBase), float64(len(inputDigits)-i-1)))
	}

	if sum == 0 {
		return []int{0}, nil
	}
	res := []int{}
	for sum > 0 {
		res = append([]int{sum % outputBase}, res...)
		sum /= outputBase
	}
	return res, nil
}
