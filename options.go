package dispatcher

// Option is a setting to change the behavior of dispatcher.
type Option func(*Config) error

// WithLogger is a Dispatcher option to set Logger.
func WithLogger(logger Logger) Option {
	return func(c *Config) error {
		if c.Logger == nil {
			c.Logger = logger
			return nil
		}
		return ErrOptionAlreadySeted

	}
}

// MaxWorkers is Dispatcher option to change the number of Dispatcher'c worker.
func MaxWorkers(n int) Option {
	return func(c *Config) error {
		if c.MaxWorkers == DefaultMaxWorkers {
			c.MaxWorkers = n
			return nil
		}
		return ErrOptionAlreadySeted
	}
}

// MaxQueues is a Dispatcher option what changes the number of wait queues for Dispatcher.
func MaxQueues(n int) Option {
	return func(c *Config) error {
		if c.MaxQueues == DefaultMaxQueues {
			c.MaxQueues = n
			return nil
		}
		return ErrOptionAlreadySeted

	}
}

// MaxRetry is a Dispatcher option what changes the number of wait queues for Dispatcher.
func MaxRetry(n int) Option {
	return func(c *Config) error {
		if c.MaxRetry == 0 {
			c.MaxRetry = n
			return nil
		}
		return ErrOptionAlreadySeted

	}
}
