package done

import "time"

type Worker struct {
	closeCh     chan struct{}
	closeDoneCh chan struct{}
}

func NewWorker() Worker {
	worker := Worker{
		closeCh:     make(chan struct{}),
		closeDoneCh: make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer func() {
			ticker.Stop()
			close(worker.closeDoneCh)
		}()

		for {
			select {
			case <-worker.closeCh:
				return
			default:
			}
			select {
			case <-worker.closeCh:
				return
			case <-ticker.C:
			}
		}

	}()

	return worker
}

func (w *Worker) ShutDown() {
	close(w.closeCh)
	<-w.closeDoneCh
}

func use2() {
	worker := NewWorker()
	time.Sleep(5 * time.Second)

	worker.ShutDown()
}
