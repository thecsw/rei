package rei

import "strconv"

// MustAtoi converts string to int, panics on error.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	Try(err)
	return i
}

// Atoi converts string to int.
func Atoi(s string) (int, error) {
	return strconv.Atoi(s)
}

// Itoa converts int to string.
func Itoa(i int) string {
	return strconv.Itoa(i)
}
