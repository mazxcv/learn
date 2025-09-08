package or

import "time"

func or[T any](chans ...<-chan T) <-chan T {
	switch len(chans) {
	case 0:
		return nil
	case 1:
		return chans[0]
	}

	done := make(chan T)

	go func() {
		defer close(done)

		switch len(chans) {
		case 2:
			select {
			case <-chans[0]:
			case <-chans[1]:
			}
		default:
			select {
			case <-chans[0]:
			case <-chans[1]:
			case <-or(chans[2:]...):
			}
		}
	}()

	return done
}

func use() {
	start := time.Now()

	<-or(
		time.After(1*time.Second),
		time.After(2*time.Second),
		time.After(3*time.Second),
	)

	println("elapsed: ", time.Since(start).Seconds())
}
