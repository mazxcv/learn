package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxWaitSecond = 2
const maxCalls = 100

// Пусть есть функция, которая выполняет какое-то долгое действие.
// Нам требуется конкурентно вызвать функцию много раз
func randomWait() int {
	workSeconds := rand.Intn(1 + maxWaitSecond)

	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

func baseConcurrency() {
	start := time.Now()
	locker := &sync.Mutex{}
	totalWorkSeconds := 0

	wg := &sync.WaitGroup{}
	wg.Add(maxCalls)
	for range maxCalls {
		go func() {
			defer wg.Done()
			seconds := randomWait()

			locker.Lock()
			totalWorkSeconds += seconds
			locker.Unlock()
		}()
	}

	wg.Wait()

	mainSeconds := time.Since(start).Seconds()
	fmt.Println("main:", mainSeconds)
	fmt.Println("total:", totalWorkSeconds)
}

func channelConcurrency() {
	start := time.Now()
	totalWorkSeconds := 0

	ch := make(chan int)

	for range maxCalls {
		go func() {
			seconds := randomWait()
			ch <- seconds
		}()
	}

	for range maxCalls {
		totalWorkSeconds += <-ch
	}

	mainSeconds := time.Since(start).Seconds()
	fmt.Println("main:", mainSeconds)
	fmt.Println("total:", totalWorkSeconds)
}
