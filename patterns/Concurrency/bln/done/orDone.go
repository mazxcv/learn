package done

import "time"

func OrDone[T any](inputCh <-chan T, doneCh <-chan struct{}) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case <-doneCh:
				return
			default:
			}
			select {
			case v, ok := <-inputCh:
				if !ok {
					return
				}
				out <- v
			case <-doneCh:
				return
			}
		}
	}()

	return out
}

func use3() {
	channel := make(chan string)

	go func() {
		for {
			channel <- "value"
			time.Sleep(1 * time.Second)
		}
	}()

	done := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	for v := range OrDone(channel, done) {
		println(v)
	}
}
