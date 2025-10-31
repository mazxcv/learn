package fanin

import "sync"

func MergeChannels[T any](channels ...<-chan T) <-chan T {

	var wg sync.WaitGroup
	wg.Add(len(channels))
	out := make(chan T)

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func use() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		for i := 0; i < 100; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()

	for v := range MergeChannels(ch1, ch2, ch3) {
		println(v)
	}
}
