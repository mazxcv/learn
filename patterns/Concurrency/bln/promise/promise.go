package promise

import "time"

type result[T any] struct {
	v   T
	err error
}

type Promise[T any] struct {
	ch chan result[T]
}

func NewPromise[T any](f func() (T, error)) *Promise[T] {
	promise := &Promise[T]{
		ch: make(chan result[T]),
	}

	go func() {
		defer close(promise.ch)

		v, err := f()
		promise.ch <- result[T]{
			v:   v,
			err: err,
		}
	}()

	return promise
}

func (p *Promise[T]) Then(onSuccess func(T), onError func(error)) {
	for r := range p.ch {
		if r.err != nil {
			onError(r.err)
		} else {
			onSuccess(r.v)
		}
	}
}

func use() {
	asyncJob := func() (string, error) {
		time.Sleep(time.Millisecond * 100)

		return "ok", nil
	}

	promise := NewPromise(asyncJob)
	promise.Then(
		func(result string) {
			println(result)
		},
		func(err error) {
			println(err)
		},
	)
}
