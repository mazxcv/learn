package pipeline

import (
	"fmt"
	"sync"
)

func split[T any](inputCh <-chan T, n int) []<-chan T {

	out := make([]chan T, n)
	for i := range n {
		out[i] = make(chan T)
	}

	go func() {
		idx := 0
		for v := range inputCh {
			out[idx] <- v
			idx = (idx + 1) % n
		}

		for _, ch := range out {
			close(ch)
		}
	}()

	res := make([]<-chan T, n)
	for i := range n {
		res[i] = out[i]
	}

	return res
}

func send3(inputCh <-chan string, n int) <-chan string {

	var wg sync.WaitGroup
	wg.Add(n)

	splitChans := split(inputCh, n)
	out := make(chan string)

	for i := range n {
		go func(idx int) {
			defer wg.Done()
			for v := range splitChans[idx] {
				out <- fmt.Sprintf("sent: %s", v)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func use3() {
	channel := make(chan string)

	go func() {
		defer close(channel)

		for range 10 {
			channel <- "value"
		}
	}()

	for v := range send3(parse(channel), 2) {
		println(v)
	}

}
