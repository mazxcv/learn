package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// go run -race main.go
func main() {
	alreadyStored := make(map[int]struct{})
	mu := sync.Mutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10))
	}

	uniqueIds := make(chan int, capacity)
	wg := sync.WaitGroup{}

	wg.Add(capacity)
	for i := 0; i < capacity; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if _, ok := alreadyStored[doubles[i]]; !ok {
				alreadyStored[doubles[i]] = struct{}{}
				uniqueIds <- doubles[i]
			}
		}()
	}

	wg.Wait()
	close(uniqueIds)
	for val := range uniqueIds {
		println(val)
	}
	fmt.Println(uniqueIds)
}
