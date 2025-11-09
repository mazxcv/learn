package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type myMutex struct {
	locked int64
}

func (m *myMutex) Lock() {
	for {
		if atomic.CompareAndSwapInt64(&m.locked, 0, 1) {
			return
		}
	}
}

func (m *myMutex) Unlock() {
	atomic.StoreInt64(&m.locked, 0)
}

func main() {
	c := 0
	wg := &sync.WaitGroup{}
	mu := myMutex{}

	for range 1000 {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			c++
		})
	}

	wg.Wait()

	fmt.Println(c)
}
