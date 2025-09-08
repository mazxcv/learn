package future

import "time"

type Future2[T any] struct {
	resCh <-chan T
}

func NewFuture2[T any](resCh <-chan T) *Future2[T] {
	return &Future2[T]{
		resCh: resCh,
	}
}

func (f *Future2[T]) Get() T {
	return <-f.resCh
}

type Promise[T any] struct {
	resCh chan T
}

func NewPromise[T any]() *Promise[T] {
	return &Promise[T]{
		resCh: make(chan T),
	}
}

func (p *Promise[T]) Set(v T) {
	p.resCh <- v
	close(p.resCh)
}

func (p *Promise[T]) GetFuture() *Future2[T] {
	return NewFuture2(p.resCh)
}

func use2() {
	promise := NewPromise[string]()

	go func() {
		time.Sleep(1 * time.Second)
		promise.Set("agreement")
	}()

	future := promise.GetFuture()

	v := future.Get()
	println(v)
}
