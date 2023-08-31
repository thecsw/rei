package rei

// Drain is a helper function that drains a channel. It is meant to be
// used as a goroutine.
func Drain[T any](vv chan T) {
	for v := range vv {
		Noop(v)
	}
}

// Noop is a helper function that does nothing.
func Noop[T any](v T) {}

// Collect returns all channel elements in a slice.
func Collect[T any](vv chan T) []T {
	ret := make([]T, 0, min(cap(vv), 1))
	for v := range vv {
		ret = append(ret, v)
	}
	return ret
}
