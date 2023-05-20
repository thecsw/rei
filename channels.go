package rei

// drain is a helper function that drains a channel. It is meant to be
// used as a goroutine.
func Drain[T any](vv chan T) {
	for v := range vv {
		Noop(v)
	}
}

// noop is a helper function that does nothing.
func Noop[T any](v T) {}
