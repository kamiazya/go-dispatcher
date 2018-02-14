package dispatcher

import "testing"

func Test_New(t *testing.T) {
	t.Parallel()
	type test struct {
		opts      []Option
		wantedErr error
	}

	testCases := []test{
		{
			opts: []Option{
				MaxQueues(1),
				MaxRetry(2),
				MaxWorkers(3),
				WithLogger(&testLogger{}),
			},
			wantedErr: nil,
		},
		{
			opts: []Option{
				MaxRetry(1),
				MaxRetry(3),
			},
			wantedErr: ErrOptionAlreadySeted,
		},
		{
			opts: []Option{
				MaxWorkers(1),
				MaxWorkers(3),
			},
			wantedErr: ErrOptionAlreadySeted,
		},
		{
			opts: []Option{
				MaxQueues(1),
				MaxQueues(3),
			},
			wantedErr: ErrOptionAlreadySeted,
		},
		{
			opts: []Option{
				WithLogger(&testLogger{}),
				WithLogger(&testLogger{}),
			},
			wantedErr: ErrOptionAlreadySeted,
		},
	}
	for _, testCase := range testCases {
		if _, err := New(testCase.opts...); err != testCase.wantedErr {
			t.Error()
		}
	}
}

func TestDispatcher(t *testing.T) {
	t.Parallel()

	t.Run("CreateWorker", func(t *testing.T) {
		// create dispatcher
		d := &dispatcher{}
		for i := 1; 10 >= i; i++ {
			// create worker from dispatcher
			w := d.createWorker()
			if int(w.id) != i {
				t.Error("worker id was not incremented. want:", i, "id:", w.id)
			}
		}
	})

}

func TestDispatcher_Dispatche(t *testing.T) {
	t.Parallel()
	testTask := func() error {
		return nil
	}

	type test struct {
		queue     int
		start     bool
		wantedErr error
	}

	testCases := []test{
		{
			// Dispatch a task to a Dispatcher that has not been started
			queue:     1,
			start:     false,
			wantedErr: nil,
		},
		{
			// Dispatch the task to the Dispatcher being started
			queue:     1,
			start:     true,
			wantedErr: nil,
		},
		{
			// Dispatch the task to a Dispatcher that has not been queued and has not been started
			queue:     0,
			start:     false,
			wantedErr: ErrQueuesCap,
		},
		{
			// Dispatch the task to the Dispatcher being started with no queue buffer
			queue:     0,
			start:     true,
			wantedErr: nil,
		},
	}

	for _, testCase := range testCases {

		d, _ := New(MaxQueues(testCase.queue))
		if testCase.start {
			d.Start()
		}
		if err := d.Dispatche(testTask); err != testCase.wantedErr {
			t.Error(err)
			t.Log(testCase)
		}
	}

}

func TestDispatcher_Start(t *testing.T) {
	t.Parallel()
	type test struct {
		start     bool
		wantedErr error
	}

	testCases := []test{
		{
			start:     false,
			wantedErr: nil,
		},
		{
			start:     true,
			wantedErr: ErrAlreadyStarted,
		},
	}

	for _, testCase := range testCases {
		d, _ := New()
		if testCase.start {
			d.Start()
		}
		if err := d.Start(); err != testCase.wantedErr {
			t.Error()
		}
	}
}
func TestDispatcher_Wait(t *testing.T) {
	t.Parallel()

	type test struct {
		start     bool
		stop      bool
		wantedErr error
	}

	testCases := []test{
		{
			start:     false,
			wantedErr: ErrNotStartedYet,
		},
		{
			start:     true,
			wantedErr: nil,
		},
		{
			start:     true,
			stop:      true,
			wantedErr: ErrAlreadyStoped,
		},
	}
	for _, testCase := range testCases {
		d, _ := New(WithLogger(&testLogger{}))
		if testCase.start {
			d.Start()
		}
		if testCase.stop {
			d.Stop(false)
		}
		if err := d.Wait(); err != testCase.wantedErr {
			t.Error()
		}
	}
}

func TestDispatcher_Stop(t *testing.T) {
	t.Parallel()

	type test struct {
		start     bool
		stop      bool
		force     bool
		wantedErr error
	}

	testCases := []test{
		{
			start:     false,
			force:     false,
			wantedErr: ErrNotStartedYet,
		},
		{
			start:     false,
			force:     true,
			wantedErr: ErrNotStartedYet,
		},
		{
			start:     true,
			force:     false,
			wantedErr: nil,
		},
		{
			start:     true,
			force:     true,
			wantedErr: nil,
		},
		{
			start:     true,
			stop:      true,
			force:     true,
			wantedErr: ErrAlreadyStoped,
		},
	}
	for _, testCase := range testCases {
		d, _ := New(WithLogger(&testLogger{}))
		if testCase.start {
			d.Start()
		}
		if testCase.stop {
			d.Stop(false)
		}
		if err := d.Stop(testCase.force); err != testCase.wantedErr {
			t.Error(err, testCase)
		}
	}
}
func TestDispatcher_Clone(t *testing.T) {
	t.Parallel()

	type test struct {
		maxWorkers int
		maxQueues  int
		maxRetry   int
		logger     Logger
	}
	testCases := []test{
		{1, 2, 3, nil},
		{1, 2, 3, &testLogger{}},
		{2, 3, 4, nil},
		{2, 3, 4, &testLogger{}},
		{10, 10, 0, nil},
		{10, 10, 0, &testLogger{}},
	}
	for _, testCase := range testCases {
		a, _ := New(
			MaxWorkers(testCase.maxWorkers),
			MaxQueues(testCase.maxQueues),
			MaxRetry(testCase.maxRetry),
			WithLogger(testCase.logger),
		)

		aInternal := a.(*dispatcher)

		// make clone
		b := a.Clone()
		bInternal := b.(*dispatcher)

		if aInternal.config.MaxWorkers != bInternal.config.MaxWorkers {
			t.Error()
		}
		if aInternal.config.MaxQueues != bInternal.config.MaxQueues {
			t.Error()
		}
		if aInternal.config.MaxRetry != bInternal.config.MaxRetry {
			t.Error()
		}
		if aInternal.config.Logger != bInternal.config.Logger {
			t.Error()
		}
	}

}
