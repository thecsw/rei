package rei

// Try panics if err is not nil.
func Try(err error) {
	if err != nil {
		panic(err)
	}
}

// Must panics if err is not nil, otherwise returns a.
func Must[T any](a T, err error) T {
	Try(err)
	return a
}

// Must2 panics if err is not nil, otherwise returns a and b.
func Must2[T1 any, T2 any](a T1, b T2, err error) (T1, T2) {
	Try(err)
	return a, b
}

// Must3 panics if err is not nil, otherwise returns a, b and c.
func Must3[T1 any, T2 any, T3 any](a T1, b T2, c T3, err error) (T1, T2, T3) {
	Try(err)
	return a, b, c
}

// Must4 panics if err is not nil, otherwise returns a, b, c and d.
func Must4[T1 any, T2 any, T3 any, T4 any](a T1, b T2, c T3, d T4, err error) (T1, T2, T3, T4) {
	Try(err)
	return a, b, c, d
}

// Must5 panics if err is not nil, otherwise returns a, b, c, d and e.
func Must5[T1 any, T2 any, T3 any, T4 any, T5 any](a T1, b T2, c T3, d T4, e T5, err error) (T1, T2, T3, T4, T5) {
	Try(err)
	return a, b, c, d, e
}
