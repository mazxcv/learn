package errorgroup

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type ErrorGroup struct {
	err  error
	wg   sync.WaitGroup
	once sync.Once

	doneCh chan struct{}
}

func NewErrorGroup() (*ErrorGroup, <-chan struct{}) {
	doneCh := make(chan struct{})

	return &ErrorGroup{
		doneCh: doneCh,
	}, doneCh
}

func (e *ErrorGroup) Go(task func() error) {
	e.wg.Add(1)

	go func() {
		defer e.wg.Done()

		select {
		case <-e.doneCh:
			return
		default:
			if err := task(); err != nil {
				e.once.Do(func() {
					e.err = err
					close(e.doneCh)
				})
			}
		}
	}()
}

func (e *ErrorGroup) Wait() error {
	e.wg.Wait()
	return e.err
}

func use() {
	group, groupDone := NewErrorGroup()

	for range 10 {
		group.Go(func() error {
			timeout := time.Second * time.Duration(rand.Intn(10))
			timer := time.NewTicker(timeout)
			defer timer.Stop()

			select {
			case <-groupDone:
				println("cancelled")
				return errors.New("cancelled")
			case <-timer.C:
				println("timeout")
				return errors.New("timeout")
			}
		})
	}

	if err := group.Wait(); err != nil {
		println(err.Error())
	}
}
