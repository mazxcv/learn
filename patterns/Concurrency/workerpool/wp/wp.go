package wp

type Worker struct {
	id            int
	jobsCompleted int
}

var workers = []*Worker{
	{id: 1},
	{id: 2},
	{id: 3},
	{id: 4},
	{id: 5},
	{id: 6},
	{id: 7},
	{id: 8},
	{id: 9},
	{id: 10},
}

type Pool[Data any] struct {
	pool    chan *Worker
	handler func(int, Data)
}

func New[Data any](handler func(int, Data)) *Pool[Data] {
	return &Pool[Data]{
		handler: handler,
		pool:    make(chan *Worker, len(workers)),
	}
}

func (p *Pool[Data]) Create() {
	for _, w := range workers {
		p.pool <- w
	}
}

func (p *Pool[Data]) Handle(data Data) {
	w := <-p.pool
	go func() {
		p.handler(w.id, data)
		p.pool <- w
		w.jobsCompleted++
	}()
}

func (p *Pool[Data]) Wait() {
	for range workers {
		<-p.pool
	}
}

func (p *Pool[Data]) Stats() {
	println("Jobs completed: ")
	for _, w := range workers {
		println("Worker", w.id, "completed", w.jobsCompleted, "jobs")
	}
}
