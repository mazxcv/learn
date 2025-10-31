package rate

import "time"

type RateLimiter struct {
	leakyBucketCh chan struct{}

	closeCh     chan struct{}
	closeDoneCh chan struct{}
}

func NewLeakyBucketLimiter(limit int, period time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		leakyBucketCh: make(chan struct{}, limit),
		closeCh:       make(chan struct{}),
		closeDoneCh:   make(chan struct{}),
	}

	leakinterval := period.Nanoseconds() / int64(limit)

	go limiter.startPeriodicLeak(time.Duration(leakinterval))

	return limiter
}

func (r *RateLimiter) startPeriodicLeak(period time.Duration) {
	timer := time.NewTicker(period)

	defer func() {
		timer.Stop()
		close(r.closeDoneCh)
	}()

	for {
		select {
		case <-r.closeCh:
			return
		default:
		}

		select {
		case <-r.closeCh:
			return
		case <-timer.C:
			select {
			case <-r.leakyBucketCh:
			default:
			}
		}
	}
}

func (r *RateLimiter) Allow() bool {
	select {
	case r.leakyBucketCh <- struct{}{}:
		return true
	default:
		return false
	}
}

func (r *RateLimiter) Shutdown() {
	close(r.closeCh)
	<-r.closeDoneCh
}
