package fanout

import "sync"

func SplitChannel[T any](inputChan <-chan T, n int) []<-chan T {
	out := make([]chan T, n)
	for i := range n {
		out[i] = make(chan T)
	}

	go func() {
		idx := 0
		for v := range inputChan {
			out[idx] <- v
			idx = (idx + 1) % len(out)
		}

		for _, ch := range out {
			close(ch)
		}
	}()

	res := make([]<-chan T, len(out))
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

	channels := SplitChannel(channel, 2)
	var wg sync.WaitGroup
	wg.Add(2)

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

	go func() {
		defer wg.Done()

		for v := range channels[1] {
			println("ch[0]:", v)
		}
	}()

	wg.Wait()

}
