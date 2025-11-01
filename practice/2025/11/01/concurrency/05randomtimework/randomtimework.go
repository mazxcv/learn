package main

import (
	"math/rand"
	"time"
)

func randomtimework() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
}

// Написать обертку для этой функции, которая прервет выполнение,
// если функция отрабатывает больше 3 секунд и вернет ошибку
func predictableWork(n int) {
	ch := make(chan struct{})

	go func() {
		defer close(ch)
		randomtimework()
	}()

	select {
	case <-ch:
	case <-time.After(time.Second * time.Duration(n)):
	}
}

func main() {
	predictableWork(3)
}
