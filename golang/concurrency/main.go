package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processData(val int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return val * 2
}

//	Надо паралельно запустить 5 воркеров
//
// Читать из in
// Писать в out
// При этом операция не должна выполняться более 5 секунд
func processParallel(in <-chan int, out chan<- int, numWorkers int) {

}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
	}()

	now := time.Now()
	processParallel(in, out, 5)

	for val := range out {
		fmt.Println(val, "in", time.Since(now))
	}

	fmt.Println(time.Since(now).Seconds())

}
