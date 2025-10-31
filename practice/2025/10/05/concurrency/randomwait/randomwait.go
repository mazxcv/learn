package randomwait

import (
	"math/rand"
	"sync"
	"time"
)

func randomwait() int {
	workSeconds := rand.Intn(5 + 1)

	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

// Требуется 100 раз вызвать функцию randomwait.,
// Но конкурентно
// Вывести общее время работы

func rnd() (int, int) {
	totalWorkSecons := 0
	start := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(100)

	lock := &sync.Mutex{}

	for range 100 {
		go func() {
			defer wg.Done()
			seconds := randomwait()
			lock.Lock()
			totalWorkSecons += seconds
			lock.Unlock()
		}()
	}

	wg.Wait()

	return int(time.Since(start).Seconds()), totalWorkSecons
}

func chRnd() (int, int) {

	ch := make(chan int)
	total := 0
	start := time.Now()

	for range 100 {
		go func() {
			seconds := randomwait()
			ch <- seconds
		}()
	}

	for range 100 {
		total += <-ch
	}

	return int(time.Since(start).Seconds()), total
}
