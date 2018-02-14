package dispatcher

// LifeCycle of Dispatcher.
//
//
// Initial state
//
//      none               New(), GenerateFromConfig(Config)
//        |
//   [ Initial ]           .Dispatch(Task)             # if dispatcher max queue is larger then zero
//        |--- .Start()
//   [ Started ]           .Wait(), .Dispatch(Task)
//        |--- .Stop()
//   [ Stoped  ]
type LifeCycle int8

//go:generate stringer -type=LifeCycle

const (
	// Initial is a state of Dispatcher that prepared to start.
	// Warker was created and are not waiting for tasks.
	//
	// You can dispatch Task even in the Initial state,
	// but if you try to dispatch a task that exceeds the Queue's buffer size limit
	// it will return ErrQueuesCap.
	Initial LifeCycle = iota

	// Started is a state of Dispatcher that is waiting for task to excute.
	Started

	// Stoped is a state of Dispatcher that is NOT waiting for task to excute.
	// If you dispatch task to that, you'll get ErrAlreadyStoped.
	//
	// To dispatch more tasks, execute .Clone () to get the Dispatcher
	// with the same configuration as the Clone source in the initial state.
	Stoped
)
