package dispatcher

import (
	"errors"
	"testing"
)

var (
	errFlg        bool
	tryCount      int
	errTestWorker = errors.New("test")
)

var _ Logger = (*testLogger)(nil)

type workerTestLogger struct{}

func (l *workerTestLogger) Log(log string) {
}

func (l *workerTestLogger) Error(err error) {
	errFlg = true
}

func taskCommon() {
	tryCount++
}

func reset() {
	errFlg = false
	tryCount = 0
}

func TestWorker(t *testing.T) {
	t.Parallel()
	type test struct {
		dispatcherFactory func() Dispatcher
		task              Task
		wantedErrFlg      bool
		wantedTryCount    int
	}

	testCases := []test{
		{
			dispatcherFactory: func() Dispatcher {
				d, _ := New(
					WithLogger(&workerTestLogger{}),
				)
				return d
			},
			task: func() error {
				return nil
			},
			wantedErrFlg:   false,
			wantedTryCount: 0,
		},
		{
			dispatcherFactory: func() Dispatcher {
				d, _ := New(
					WithLogger(&workerTestLogger{}),
					MaxRetry(5),
				)
				return d
			},
			task: func() error {
				taskCommon()
				if 3 > tryCount {
					return errTestWorker
				}
				return nil
			},
			wantedErrFlg:   true,
			wantedTryCount: 3,
		},
		{
			dispatcherFactory: func() Dispatcher {
				d, _ := New(
					WithLogger(&workerTestLogger{}),
					MaxRetry(5),
				)
				return d
			},
			task: func() error {
				taskCommon()
				if 10 > tryCount {
					return errTestWorker
				}
				return nil
			},
			wantedErrFlg:   true,
			wantedTryCount: 6,
		},
	}

	for _, testCase := range testCases {
		d := testCase.dispatcherFactory()
		d.Start()

		d.Dispatche(testCase.task)

		d.Wait()

		if testCase.wantedErrFlg != errFlg {
			t.Error(testCase.wantedErrFlg, errFlg)
		}

		if testCase.wantedTryCount != tryCount {
			t.Error(testCase.wantedTryCount, tryCount)
		}

		reset()
	}

}
