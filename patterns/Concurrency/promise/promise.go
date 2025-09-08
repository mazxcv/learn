package promise

import "time"

type result[T any] struct {
	v   T
	err error
}

type Promise[T any] struct {
	resultCh chan result[T]
}

func NewPromise[T any](asyncF func() (T, error)) Promise[T] {
	promise := Promise[T]{
		resultCh: make(chan result[T]),
	}

	go func() {
		defer close(promise.resultCh)

		v, err := asyncF()
		promise.resultCh <- result[T]{v: v, err: err}
	}()

	return promise
}

func (p *Promise[T]) Then(successF func(T), errorF func(error)) {
	go func() {
		res := <-p.resultCh

		if res.err != nil {
			errorF(res.err)
		} else {
			successF(res.v)
		}

	}()
}

func use() {
	asyncJob := func() (string, error) {
		time.Sleep(1 * time.Second)

		return "ok", nil
	}

	promise := NewPromise(asyncJob)
	promise.Then(
		func(v string) {
			println("Success:", v)
		},
		func(err error) {
			println("Error:", err.Error())
		},
	)

	time.Sleep(2 * time.Second)
}
