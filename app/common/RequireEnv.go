package common

import (
	"log"
	"os"
	"strings"
)

// RequireEnv - ensure that we don't allow empty strings or leading/trailing whitespace
func RequireEnv(key string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		log.Fatalf("missing environment variable: %s", key)
	}
	return val
}
