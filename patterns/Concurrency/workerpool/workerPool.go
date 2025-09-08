package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func Start[T, R any](ctx context.Context, workerCount int, input <-chan T, f func(t T) R) <-chan R {
	out := make(chan R)
	wg := &sync.WaitGroup{}

	for range workerCount {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case t, ok := <-input:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- f(t):
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Generator[T any](ctx context.Context, data []T, size int) <-chan T {
	out := make(chan T, size)

	defer func() {
		defer close(out)
		for _, v := range data {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}()

	return out
}

func use() {

	ctx := context.Background()

	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.somestupidname.com/",
		"http://www.nothing.com/",
	}

	type result struct {
		url string
		err error
	}

	inputChans := Generator(ctx, urls, len(urls))
	resChans := Start(ctx, 10, inputChans, func(url string) result {
		resp, err := http.Get(url)
		if err != nil {
			return result{url: url, err: err}
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return result{url: url, err: fmt.Errorf("status code: %d", resp.StatusCode)}
		}

		return result{url: url}
	})

	for res := range resChans {
		println(res.url, res.err)
	}

}
