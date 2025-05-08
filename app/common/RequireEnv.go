package common

import (
	"log"
	"os"
	"strings"
)

func RequireEnv(key string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		log.Fatalf("missing environment variable: %s", key)
	}
	return val
}
