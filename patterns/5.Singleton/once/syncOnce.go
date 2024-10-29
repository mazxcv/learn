package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("In Thread: Single instance already created")
	}

	return singleInstance
}
