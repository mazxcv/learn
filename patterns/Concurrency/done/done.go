package done

func process(doneCh <-chan struct{}) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			select {
			case <-doneCh:
				return
			default:
			}
		}
	}()

	return done
}

func use() {
	closeCh := make(chan struct{})
	closeDoneCh := process(closeCh)

	close(closeCh)
	<-closeDoneCh

	println("terminated")
}
