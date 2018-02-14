package dispatcher

// Dispatcher represents a management workers.
type Dispatcher interface {
	// Dispatche adds a given task to the queue of the dispatcher.
	Dispatche(t Task) error

	// Start starts the specified dispatcher
	// but does not wait for it to complete.
	Start() error

	// Wait waits for the dispatcher to exit.
	// It must have been started by Start.
	Wait() error

	// Stop stops the dispatcher to execute.
	// The dispatcher stops gracefully
	// if the given boolean is false.
	Stop(force bool) error

	// Clone creates Dispatcher of the same configuration
	// from Config when Dispatcher was created.
	Clone() Dispatcher

	// LifeCycle returns LifeCycle of Dispatcher.
	LifeCycle() LifeCycle
}

// New returns Dispatcher.
// You can set it some options.
func New(opts ...Option) (Dispatcher, error) {
	// Get default setting
	config := DafaultConfig()

	// Apply custom options to setting if exists
	if len(opts) > 0 {
		for _, opt := range opts {
			if err := opt(config); err != nil {
				// if error occurred return error
				return nil, err
			}
		}
	}

	// Generate Dispatcher from Config.
	return GenerateFromConfig(*config)
}
