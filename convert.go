package rei

import "strconv"

// atoi converts string to int.
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	try(err)
	return i
}

// itoa converts int to string.
func itoa(i int) string {
	return strconv.Itoa(i)
}
