package grains

import (
	"fmt"
)

func Square(number int) (uint64, error) {
	if number > 64 || number <= 0 {
		return 0, fmt.Errorf("64 squares on a chessboard (")
	}
	return 1 << uint(number-1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
