package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomwait() int {
	randomWaitSecons := rand.Intn(5 + 1)

	time.Sleep(time.Second * time.Duration(randomWaitSecons))

	return randomWaitSecons
}

func main() {
	start := time.Now()
	totalTimeWork := 0

	ch := make(chan int)

	for range 100 {
		go func() {
			seconds := randomwait()
			ch <- seconds
		}()
	}

	for range 100 {
		totalTimeWork += <-ch
	}

	fmt.Printf("total time: %d, work time: %v seconds\n", totalTimeWork, time.Since(start).Seconds())
}
