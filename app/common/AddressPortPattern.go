package common

import (
	"log"
	"net"
	"strconv"
	"strings"
)

func AddressPortPattern(input string) string {
	host, portStr, err := net.SplitHostPort(input)
	if err != nil || strings.TrimSpace(host) == "" {
		log.Fatalf("invalid address:port format: %s", input)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		log.Fatalf("invalid port number: %s", portStr)
	}

	return input
}
