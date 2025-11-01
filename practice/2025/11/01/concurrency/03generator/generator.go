package main

import (
	"fmt"
	"sync"
)

func generator(count int) <-chan int {
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(count)

	for i := range count {
		go func(i int) {
			defer wg.Done()

			for j := range 5 {
				ch <- 100*i + j + 1
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	ch := generator(5)

	for v := range ch {
		fmt.Println(v)
	}
}
