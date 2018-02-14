package dispatcher_test

import (
	"fmt"

	"github.com/kamiazya/go-dispatcher"
)

// make sure to sampleDefaultLogger implements Logger iterface.
var _ dispatcher.Logger = (*sampleDefaultLogger)(nil)

type sampleDefaultLogger struct {
	// implements Log and Error
	sampleLogger
}

// String for fmt.Stringer
func (l *sampleDefaultLogger) String() string {
	return "sample default logger"
}

func Example_variables() {
	dispatcher.DefaultMaxWorkers = 1
	dispatcher.DefaultMaxQueues = 2
	dispatcher.DefaultMaxRetry = 3
	dispatcher.DafaultLogger = &sampleDefaultLogger{}

	// check default variables.
	fmt.Println("DefaultMaxWorkers", dispatcher.DefaultMaxWorkers)
	fmt.Println("DefaultMaxQueues", dispatcher.DefaultMaxQueues)
	fmt.Println("DefaultMaxRetry", dispatcher.DefaultMaxRetry)
	fmt.Println("DafaultLogger", dispatcher.DafaultLogger)

	// check default configs.
	config := dispatcher.DafaultConfig()
	fmt.Println("config.MaxWorkers", config.MaxWorkers)
	fmt.Println("config.MaxQueues", config.MaxQueues)
	fmt.Println("config.MaxRetry", config.MaxRetry)
	fmt.Println("config.Logger", config.Logger)

	// Output:
	// DefaultMaxWorkers 1
	// DefaultMaxQueues 2
	// DefaultMaxRetry 3
	// DafaultLogger sample default logger
	// config.MaxWorkers 1
	// config.MaxQueues 2
	// config.MaxRetry 3
	// config.Logger sample default logger
}
