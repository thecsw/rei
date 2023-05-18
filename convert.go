package rei

import "strconv"

// Atoi converts string to int.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Try(err)
	return i
}

// Itoa converts int to string.
func Itoa(i int) string {
	return strconv.Itoa(i)
}
