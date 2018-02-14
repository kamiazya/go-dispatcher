package dispatcher

import (
	"fmt"
)

// worker represents the worker that executes the job.
type worker struct {
	id         uint64
	dispatcher *dispatcher

	taskSeq uint64
	task    chan Task

	// quit signal receiver
	quit chan struct{}
}

func (w *worker) start() {
	go func() {
		for {
			// register the current worker into the dispatch pool
			w.dispatcher.pool <- w

			select {
			case task := <-w.task: // on receive Task

				// excute task
				w.excute(task)

			case <-w.quit: // on receive quit signal
				// log
				if w.dispatcher.hasLogger {
					w.dispatcher.logger.Log(fmt.Sprintf("dispatcher: stoped worker%d", w.id))
				}
				return
			}
		}
	}()
}

func (w *worker) excute(task Task) {
	// increment taskSeq
	w.taskSeq++

	// log
	if w.dispatcher.hasLogger {
		w.dispatcher.logger.Log(fmt.Sprintf("dispatcher: task %d excuted on worker %d", w.taskSeq, w.id))
	}

	// task excute
	if err := task(); err != nil {
		// on error

		// log
		if w.dispatcher.hasLogger {
			w.dispatcher.logger.Error(fmt.Errorf("dispatcher: task %d failed on worker%d: %s", w.taskSeq, w.id, err.Error()))
		}

		// retry
		w.retry(task)
	}

	if w.dispatcher.hasLogger {
		w.dispatcher.logger.Log(fmt.Sprintf("dispatcher: task %d done on worker %d", w.taskSeq, w.id))
	}
	w.dispatcher.wg.Done()
}

func (w *worker) retry(task Task) {
	for i := 0; w.dispatcher.config.MaxRetry > i; i++ {

		if err := task(); err != nil {
			// on error

			// log
			if w.dispatcher.hasLogger {
				w.dispatcher.logger.Error(fmt.Errorf("dispatcher: task %d failed on worker%d: retry count %d: %s", w.taskSeq, w.id, i+1, err.Error()))
			}
		} else {
			// on successed

			// log
			if w.dispatcher.hasLogger {
				w.dispatcher.logger.Log(fmt.Sprintf("dispatcher: task %d successed on worker%d: retry count %d", w.taskSeq, w.id, i+1))
			}

			// end retry loop
			break
		}
	}
}

func (w *worker) stop() {
	// log
	if w.dispatcher.hasLogger {
		w.dispatcher.logger.Log(fmt.Sprintf("dispatcher: stopping worker%d", w.id))
	}

	// send stop signal
	w.quit <- struct{}{}
}
