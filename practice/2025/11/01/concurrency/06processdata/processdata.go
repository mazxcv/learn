package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// существует функция processData, которая рандомное количество времени ждет
// и возвращает x2

// создать функцию, которая будет зпускать параллельно processData в нескольких потоках
// при этом вся операция не должна занимать более 5 секунд

type outVal struct {
	val int
	err error
}

func processData(ctx context.Context, v int) chan outVal {
	ch := make(chan struct{})

	out := make(chan outVal)

	go func() {
		defer close(ch)
		time.Sleep(time.Duration(rand.Intn(7)) * time.Second)
	}()

	go func() {
		select {
		case <-ctx.Done():
			out <- outVal{err: ctx.Err()}
		case <-ch:
			out <- outVal{val: v * 2}
		}
	}()

	return out
}

func main() {
	in := make(chan int)
	out := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go func() {
		defer close(in)
		for i := range 10 {
			select {
			case <-ctx.Done():
				return
			case in <- i + 1:
			}
		}
	}()

	start := time.Now()
	processparallel(ctx, in, out, 10)

	for v := range out {
		fmt.Println(v)
	}

	fmt.Println(time.Since(start))
}

func processparallel(ctx context.Context, in <-chan int, out chan<- int, cnt int) {
	wg := &sync.WaitGroup{}
	for range cnt {
		wg.Go(func() {
			worker(ctx, in, out)
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func worker(ctx context.Context, in <-chan int, out chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			select {
			case ov := <-processData(ctx, v):
				if ov.err != nil {
					return
				}
				select {
				case out <- ov.val:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}
}
