package rei

// Try panics if err is not nil.
func Try(err error) {
	if err != nil {
		panic(err)
	}
}
