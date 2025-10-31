package pipeline

func add(done <-chan struct{}, in <-chan int, delta int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			select {
			case <-done:
				return
			case out <- v + delta:
			}
		}
	}()

	return out
}

func main() {
	stream := make(chan int, 5)

	for i := range 5 {
		stream <- i
	}
	close(stream)

	done := make(chan struct{})
	defer close(done)

	pipeline := add(done, stream, 2)

	for v := range pipeline {
		println(v)
	}
}
