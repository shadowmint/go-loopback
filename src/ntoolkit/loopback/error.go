package loopback

const (
	// ErrListen means binding a random local port failed.
	ErrListen int = iota

	// ErrConnect indicates connecting to the loopback port failed.
	ErrConnect
)
