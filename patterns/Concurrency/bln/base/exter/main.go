package base

import (
	"context"
	"fmt"
)

var numWorkers = 10

func main() {
	in := make(chan int)

	go func() {
		defer close(in)
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for v := range fanIn(ctx, fanOut(in, numWorkers, processing)) {
		fmt.Println(v)
	}
}
