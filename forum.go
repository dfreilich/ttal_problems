package ttal

import (
	"strings"
)

func TrimMessage(message string, K int) string {
	if K <= 0 {
		return ""
	} else if K > len(message) {
		return strings.TrimSpace(message)
	}

	if len(message) <= K || string(message[K]) == " " {
		return strings.TrimSpace(message[:K])
	}
	lastSpace := strings.LastIndex(message[:K], " ")
	if lastSpace == -1 {
		return ""
	}
	return strings.TrimSpace(message[:lastSpace])
}
