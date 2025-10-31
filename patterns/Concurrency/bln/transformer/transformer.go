package transformer

func Transformer[T any](inputCh <-chan T, f func(T) T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range inputCh {
			out <- f(v)
		}
	}()
	return out
}

func use() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 100; i++ {
			channel <- i
		}
	}()

	mul := func(val int) int {
		return val * val
	}

	for number := range Transformer(channel, mul) {
		println(number)
	}
}
