package exc

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func exc() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		// ctx
		// compute

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return nil
		}
	})
	group.Go(compute)
	group.Go(compute)
	group.Go(compute)
	group.Go(compute)

	err := group.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func compute() error {
	sec := rand.Intn(3)
	dur := time.Duration(sec) * time.Second
	time.Sleep(dur)

	if sec == 1 {
		return errors.New("timeout")
	}

	return nil
}
