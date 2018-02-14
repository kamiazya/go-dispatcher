package dispatcher

var _ Dispatcher = (*dispatcher)(nil)

// Dispatche adds a given task to the queue of the dispatcher.
func (d *dispatcher) Dispatche(t Task) error {
	if d.lifecycle == Stoped {
		// If already stoped
		return ErrAlreadyStoped
	} else if d.lifecycle == Initial && cap(d.queue) == len(d.queue) {
		// If lifecycle is Initial, it able to dispatch task before rearch limit of task queue.
		return ErrQueuesCap
	}

	d.dispatch(t)
	return nil
}

// Start starts the specified dispatcher but does not wait for it to complete.
func (d *dispatcher) Start() error {
	switch d.lifecycle {
	case Started:
		// If already started
		return ErrAlreadyStarted
	case Stoped:
		// If already stoped
		return ErrAlreadyStoped
	}

	// set the started flag to true
	d.lifecycle = Started

	// Start all workers
	for _, w := range d.workers {
		w.start()
	}

	go d.main()
	return nil
}

// Wait waits for the dispatcher to exit. It must have been started by Start.
func (d *dispatcher) Wait() error {
	switch d.lifecycle {
	case Initial:
		// If not started yet
		return ErrNotStartedYet
	case Stoped:
		// If already stoped
		return ErrAlreadyStoped
	}
	d.wait()
	return nil
}

// Stop stops the dispatcher to execute. The dispatcher stops gracefully
// if the given boolean is false.
func (d *dispatcher) Stop(force bool) error {
	switch d.lifecycle {
	case Initial:
		// If not started yet
		return ErrNotStartedYet
	case Stoped:
		// If already stoped
		return ErrAlreadyStoped
	}
	if !force {
		// err is already checked.
		d.Wait()
	}

	d.stop()
	return nil
}

// Clone creates Dispatcher of the same configuration
// from Config when Dispatcher was created.
func (d *dispatcher) Clone() Dispatcher {
	clone, _ := GenerateFromConfig(d.config)
	return clone
}

// LifeCycle returns LifeCycle of Dispatcher.
func (d *dispatcher) LifeCycle() LifeCycle {
	return d.lifecycle
}
