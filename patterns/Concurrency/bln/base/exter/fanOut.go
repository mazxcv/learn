package base

func fanOut(ch chan int, numChans int, f func(int) int) []chan int {

	chans := make([]chan int, numChans)

	for i := range numChans {
		chans[i] = pipeline(ch, f)
	}

	return chans
}

func pipeline(in chan int, f func(int) int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- f(v)
		}
	}()

	return out
}
