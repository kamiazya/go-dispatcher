package dispatcher

// Logger is an interface provided to allow
// users to use arbitrary logging libraries.
type Logger interface {
	Log(string)
	Error(error)
}
