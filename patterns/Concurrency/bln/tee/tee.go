package tee

import "sync"

func Tee[T any](inputCh <-chan T, n int) []<-chan T {
	out := make([]chan T, n)
	for i := range n {
		out[i] = make(chan T)
	}

	go func() {
		for v := range inputCh {
			for i := range n {
				out[i] <- v
			}
		}

		for _, ch := range out {
			close(ch)
		}
	}()

	res := make([]<-chan T, n)
	for i := range out {
		res[i] = out[i]
	}
	return res
}

func use() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 100; i++ {
			channel <- i
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	channels := Tee(channel, 2)
	go func() {
		defer wg.Done()
		for v := range channels[0] {
			println(v)
		}
	}()

	go func() {
		defer wg.Done()

		for v := range channels[0] {
			println("ch[0]:", v)
		}
	}()

	wg.Wait()
}
