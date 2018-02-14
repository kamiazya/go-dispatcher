package dispatcher

var _ Logger = (*testLogger)(nil)

type testLogger struct{}

func (l *testLogger) Log(log string) {
}

func (l *testLogger) Error(err error) {
}
