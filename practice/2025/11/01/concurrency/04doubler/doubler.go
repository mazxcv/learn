package main

import (
	"fmt"
	"time"
)

// Написать 3 функции
// writer - генерирует числа от 1 до 10
// doubler - умножает на 2 и имитирует долгую работу 500 ms
// reader - читает числа

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func doubler(ch <-chan int) <-chan int {
	res := make(chan int)

	go func() {

		defer close(res)
		for v := range ch {
			time.Sleep(time.Millisecond * 500)
			res <- v * 2
		}
	}()

	return res
}

func writer() <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		i := 0
		for {
			res <- i + 1

			i++
		}
	}()

	return res
}

func main() {
	reader(doubler(writer()))
}
