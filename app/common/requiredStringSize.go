package common

import "log"

// RequiredStringSize - ensure that the input string is of at least sz characters
func RequiredStringSize(sz uint, input string) string {
	if uint(len(input)) < sz {
		log.Fatalf("input string size too small (requires %d)", sz)
	}
	return input
}
