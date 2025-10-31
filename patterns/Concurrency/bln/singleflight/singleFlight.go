package singleflight

import (
	"sync"
	"time"
)

type call struct {
	err  error
	v    any
	done chan struct{}
}

type singleFlight struct {
	mutex sync.Mutex
	calls map[string]*call
}

func NewSingleFlight() *singleFlight {
	return &singleFlight{
		calls: make(map[string]*call),
	}
}

func (s *singleFlight) Do(key string, fn func() (any, error)) (any, error) {
	s.mutex.Lock()

	if call, ok := s.calls[key]; ok {
		s.mutex.Unlock()
		return s.wait(call)
	}

	call := &call{
		done: make(chan struct{}),
	}

	s.calls[key] = call
	s.mutex.Unlock()

	go func() {
		defer func() {
			s.mutex.Lock()
			close(call.done)
			delete(s.calls, key)
			s.mutex.Unlock()
		}()
	}()

	return s.wait(call)
}

func (s *singleFlight) wait(call *call) (any, error) {
	<-call.done
	return call.v, call.err
}

func use() {
	const inFlightRecords = 100

	var wg sync.WaitGroup
	wg.Add(inFlightRecords)

	singleFlight := NewSingleFlight()

	const key = "same key"

	for i := range inFlightRecords {
		go func() {
			defer wg.Done()

			v, err := singleFlight.Do(key, func() (interface{}, error) {
				println("single flight")
				time.Sleep(5 * time.Second)

				return "result", nil
			})

			println(i, " = ", v, err)
		}()
	}

	wg.Wait()
}
