package bridge

func Bridge[T any](inputChOfCh chan chan T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for inputCh := range inputChOfCh {
			for v := range inputCh {
				out <- v
			}
		}
	}()

	return out
}

func use() {
	channelOfChannel := make(chan chan string)

	go func() {
		chan1 := make(chan string, 3)
		for i := 0; i < 3; i++ {
			chan1 <- "channel-1"
		}
		close(chan1)

		chan2 := make(chan string, 3)
		for i := 0; i < 3; i++ {
			chan2 <- "channel-2"
		}
		close(chan2)

		channelOfChannel <- chan1
		channelOfChannel <- chan2

		close(channelOfChannel)
	}()

	for v := range Bridge(channelOfChannel) {
		println(v)
	}
}
