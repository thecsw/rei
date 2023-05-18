package rei

// try panics if err is not nil.
func try(err error) {
	if err != nil {
		panic(err)
	}
}
