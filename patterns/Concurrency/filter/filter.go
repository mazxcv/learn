package filter

func Filter[T any](inputCh <-chan T, filter func(T) bool) <-chan T {
	out := make(chan T)

	go func() {
		for v := range inputCh {
			if filter(v) {
				out <- v
			}
		}
	}()

	return out
}

func use() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := range 100 {
			channel <- i
		}
	}()

	isPredicate := func(val int) bool {
		return val%2 != 0
	}

	for number := range Filter(channel, isPredicate) {
		println(number)
	}
}
