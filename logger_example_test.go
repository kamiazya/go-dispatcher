package dispatcher_test

import (
	"fmt"

	"github.com/kamiazya/go-dispatcher"
)

// make sure to sampleLogger implementes Logger iterface.
var _ dispatcher.Logger = &sampleLogger{}

type sampleLogger struct{}

func (l *sampleLogger) Log(log string) {
	fmt.Println(log)
}

func (l *sampleLogger) Error(err error) {
	fmt.Println(err.Error())
}

func ExampleLogger() {

	// create logger
	logger := new(sampleLogger)

	// create dispacher with logger
	d, err := dispatcher.New(
		dispatcher.WithLogger(logger),
		dispatcher.MaxWorkers(1),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// start dispacher
	d.Start()

	d.Dispatche(func() error {
		fmt.Println("hello")
		return nil
	})

	d.Wait()

	// Unorderd output:
	// dispatcher: Waiting for all tasks done
	// dispatcher: task 1 excuted on worker 1
	// hello
	// dispatcher: task 1 done on worker 1
	// dispatcher: All task is done
}
