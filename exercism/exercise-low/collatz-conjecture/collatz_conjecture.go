package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(n int) (i int, err error) {
	if n == 0 {
		err = fmt.Errorf("zero value")
		return
	}
	if n < 0 {
		err = fmt.Errorf("negative value")
		return
	}
	for n != 1 {
		i++
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return
}
