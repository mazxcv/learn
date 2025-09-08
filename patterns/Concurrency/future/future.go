package future

import "time"

type Future[T any] struct {
	resCh chan T
}

func NewFuture[T any](asyncF func() T) Future[T] {

	future := Future[T]{
		resCh: make(chan T),
	}

	go func() {
		defer close(future.resCh)

		future.resCh <- asyncF()
	}()

	return future
}

func (f *Future[T]) Get() T {
	return <-f.resCh
}

func use() {
	asyncJob := func() interface{} {
		time.Sleep(1 * time.Second)
		return "ok"
	}

	future := NewFuture(asyncJob)
	println(future.Get().(string))
}
