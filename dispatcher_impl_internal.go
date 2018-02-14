package dispatcher

import (
	"sync"
)

// dispatcher represents a management workers.
type dispatcher struct {
	config Config

	// worker count
	workerSeq uint64

	// internal
	pool      chan *worker
	queue     chan Task
	workers   []*worker
	wg        sync.WaitGroup
	quit      chan struct{}
	lifecycle LifeCycle
	hasLogger bool
	logger    Logger
}

func (d *dispatcher) dispatch(t Task) {
	d.wg.Add(1)
	d.queue <- t
}

func (d *dispatcher) createWorker() *worker {
	d.workerSeq++
	return &worker{
		id:         d.workerSeq,
		dispatcher: d,
		task:       make(chan Task),
		quit:       make(chan struct{}),
	}
}

func (d *dispatcher) main() {
	// main loop
	for {
		select {
		case t := <-d.queue:
			// on receive task

			// get worker from pool
			w := <-d.pool
			// dispatch task to worker
			w.task <- t

		case <-d.quit:
			// on receive quit signal
			// return to break main loop
			return
		}
	}
}

func (d *dispatcher) stop() {

	// log
	if d.hasLogger {
		d.logger.Log(msgStoppingDispatcher)
	}
	d.quit <- struct{}{}
	d.lifecycle = Stoped
	close(d.quit)
	close(d.queue)
	for _, w := range d.workers {
		w.stop()
	}
}

func (d *dispatcher) wait() {
	if d.hasLogger {
		d.logger.Log(msgWaitingForAllTasksDone)
	}

	// wait for all tasks done.
	d.wg.Wait()

	if d.hasLogger {
		d.logger.Log(msgAllTaskIsDone)
	}
}
