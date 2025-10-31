package generator

func GeneratorWothChan(start, end int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := start; i < end; i++ {
			out <- i
		}
	}()

	return out
}

func use() {
	for i := range GeneratorWothChan(100, 200) {
		println(i)
	}
}
