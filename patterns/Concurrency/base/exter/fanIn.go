package base

import (
	"context"
	"sync"
)

func fanIn(ctx context.Context, chans []chan int) chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		wg := &sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func() {
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							return
						}
						select {
						case out <- v:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
			}()
		}

		wg.Wait()
	}()

	return out
}
