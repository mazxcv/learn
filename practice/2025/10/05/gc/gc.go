// Иллюстрация работы GC в Go
package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	a []byte
}

// go:noinline
func getData() *data {
	arr := [300]byte{}
	return &data{
		a: arr[:],
	}
}

func main() {
	t1 := time.NewTicker(time.Microsecond * 100) // ~10 000 RPS

	go func() {
		for range t1.C {
			getData()
		}
	}()

	t := time.NewTicker(time.Second * 1)

	var m runtime.MemStats
	now := time.Now()

	for curr := range t.C {
		runtime.ReadMemStats(&m)

		fmt.Printf("GC_enabled: %v GC_runs: %v, live_now: %v, pause_total_ms: %.2f time %5.0f sec\n",
			m.EnableGC, m.NumGC, m.Mallocs-m.Frees, float64(m.PauseTotalNs)/1000/1000, curr.Sub(now).Seconds())
	}

}
