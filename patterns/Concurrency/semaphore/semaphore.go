package semaphore

import (
	"sync"
	"time"
)

type Semaphore struct {
	tickets chan struct{}
}

func NewSemaphore(ticketNumber int) *Semaphore {
	return &Semaphore{
		tickets: make(chan struct{}, ticketNumber),
	}
}

func (s *Semaphore) Acquire() {
	s.tickets <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.tickets
}

const countWorker = 6
const sem = 3

func use() {
	wg := sync.WaitGroup{}
	wg.Add(countWorker)
	semaphore := NewSemaphore(sem)

	for range countWorker {
		semaphore.Acquire()
		go func() {
			defer func() {
				wg.Done()
				semaphore.Release()
			}()

			println("working... ")
			time.Sleep(2 * time.Second)
			println("done")
		}()
	}

	wg.Wait()
}

// Сервис на не более чем 100 одновременных воркеров
