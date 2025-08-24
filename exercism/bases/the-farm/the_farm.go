package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodderCalculator FodderCalculator, cowCount int) (float64, error) {
	if cowCount == 0 {
		return 0, errors.New("something went wrong")
	}
	fodderAmount, err := fodderCalculator.FodderAmount(cowCount)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodderAmount * fatteningFactor / float64(cowCount), nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cowCount int) (float64, error) {
	if cowCount <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodderCalculator, cowCount)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	cowCount int
}

func (e *InvalidCowsError) Error() string {
	if e.cowCount < 0 {
		return fmt.Sprintf("%d cows are invalid: %s", e.cowCount, "there are no negative cows")
	}
	if e.cowCount == 0 {
		return fmt.Sprintf("%d cows are invalid: %s", e.cowCount, "no cows don't need food")
	}
	return ""

}

func ValidateNumberOfCows(cowCount int) error {
	if cowCount > 0 {
		return nil
	}
	return &InvalidCowsError{
		cowCount: cowCount,
	}
}
