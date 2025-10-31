package base

import "time"

func ioWait() {
	time.Sleep(time.Second * 10)
}

func cpuIntensive() {
	counter := 0

	for range 1000000000 {
		counter++
	}
}

func timeConsuming() {
	ioWait()
	// cpuIntensive()
}
