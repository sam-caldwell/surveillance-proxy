package common

import "log"

func RequiredStringSize(sz uint, input string) string {
	if uint(len(input)) < sz {
		log.Fatalf("input string size too small (requires %d)", sz)
	}
	return input
}
