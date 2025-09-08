package pipeline

func generate[T any](values ...T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for _, v := range values {
			out <- v
		}
	}()

	return out
}

func process[T any](inputCh <-chan T, f func(T) T) <-chan T {
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
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mul := func(v int) int {
		return v * v
	}

	for v := range process(generate(values...), mul) {
		println(v)
	}
}
