package lsproduct

import (
	"fmt"
	"strconv"
	"strings"
)

// func parseBigInt make array of symbol Integer from digits-string
func parseBigInt(digits string) ([]int64, error) {
	aDigits := make([]int64, 0)

	for i, v := range strings.Split(digits, "") {
		nDigits, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return []int64{}, fmt.Errorf("non Integer value")
		}
		if i == 0 && nDigits == 0 {
			continue
		}
		aDigits = append(aDigits, nDigits)
	}
	return aDigits, nil
}

// func LargestSeriesProduct
func LargestSeriesProduct(digits string, span int) (max int64, err error) {

	aDigits, err := parseBigInt(digits)
	if err != nil {
		return 0, err
	}

	if span > len(aDigits) {
		return 0, fmt.Errorf("span value is over position digits")
	}

	if span < 0 {
		return 0, fmt.Errorf("span must not be negative")
	}
	if span == 0 {
		return 0, fmt.Errorf("span must not be zero value")
	}

	for i := 0; i <= len(aDigits)-span; i++ {
		var sum int64 = 1
		for j := i; j < i+span; j++ {
			sum *= aDigits[j]
		}
		if max <= sum {
			max = sum
		}
	}
	return max, nil
}
