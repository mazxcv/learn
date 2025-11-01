package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomwait() int {
	workSeconds := rand.Intn(5 + 1)

	time.Sleep(time.Second * time.Duration(workSeconds))

	return workSeconds
}

func main() {

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	start := time.Now()
	totalWorkSeconds := 0

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workTime := randomwait()
			mu.Lock()
			totalWorkSeconds += workTime
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Printf("total time: %d, workTime: %v seconds", totalWorkSeconds, time.Since(start).Seconds())
}
