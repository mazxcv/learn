package pipeline

import (
	"fmt"
	"sync"
)

func parse(inputCh <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for data := range inputCh {
			out <- fmt.Sprintf("parsed: %s", data)
		}
	}()

	return out
}

func send(inputCh <-chan string, n int) <-chan string {
	out := make(chan string)

	var wg sync.WaitGroup
	wg.Add(n)

	for range n {
		go func() {
			defer wg.Done()

			for v := range inputCh {
				out <- fmt.Sprintf("sent: %s", v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func use2() {
	channel := make(chan string)

	go func() {
		defer close(channel)

		for range 10 {
			channel <- "value"
		}
	}()

	for v := range send(parse(channel), 2) {
		println(v)
	}

}
