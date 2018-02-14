package dispatcher

var _ Logger = &testLogger{}

type testLogger struct{}

func (l *testLogger) Log(log string) {
}

func (l *testLogger) Error(err error) {
}
