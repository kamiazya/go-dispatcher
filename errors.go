package dispatcher

import "errors"

// errors
var (
	// ErrOptionAlreadySeted if it is already set to that value.
	ErrOptionAlreadySeted = errors.New("dispatcher: The already set value")

	// ErrInvalidSetting is an error that occurs when setting an option was invalid.
	ErrInvalidSetting = errors.New("dispatcher: invalid value was seted")

	// ErrAlreadyStarted is an error that occurs when the already started Dispatcher is restarted.
	ErrAlreadyStarted = errors.New("dispatcher: already started")

	// ErrAlreadyStoped is an error that occurs when the already stoped Dispatcher is restarted.
	ErrAlreadyStoped = errors.New("dispatcher: already stoped")

	// ErrNotStartedYet is an error that occurs when a Dispatcher that has not started has been waited or stopped.
	ErrNotStartedYet = errors.New("dispatcher: not started yet")

	// ErrQueuesCap is an error that occurs when the Task is dispatched to a Dispatcher queue buffer size over queue buffer cap.
	ErrQueuesCap = errors.New("dispatcher: no buffer for dispatching task")

	// ErrFormatterIsNotAbleToBeNil is an error that occurs when set to nil ErrFormatter.
	ErrErrorFormatterIsNotAbleToBeNil = errors.New("dispatcher: don't set nil ErrFormatter")
)
