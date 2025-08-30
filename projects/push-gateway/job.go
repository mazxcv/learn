package main

import (
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	registry   = prometheus.NewRegistry()
	counterOne = promauto.With(registry).NewCounter(prometheus.CounterOpts{
		Name: "counter_one",
		Help: "Counter one help",
	})

	counterTwo = promauto.With(registry).NewCounter(prometheus.CounterOpts{
		Name: "counter_two",
		Help: "Counter two help",
	})
)

func main() {
	counterOne.Add(rand.Float64())
	counterTwo.Add(rand.Float64())

	if err := push.New("localhost:19091", "myJob").Gatherer(registry).Push(); err != nil {
		panic(err)
	}
}
