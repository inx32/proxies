package proxies

type ErrorInfo struct {
	message string
	info    *any
}

func (e *ErrorInfo) Error() string {
	return e.message
}

func (e *ErrorInfo) Info() *any {
	return e.info
}

// custom error types

type InvalidUsername struct{ ErrorInfo }
type InvalidScheme struct{ ErrorInfo }
type InvalidPort struct{ ErrorInfo }
type InvalidURL struct{ ErrorInfo }
type DialerInitFailed struct{ ErrorInfo }
type ConnectionCheckFailed struct{ ErrorInfo }
