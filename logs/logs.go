package logs

import (
	"log"
	"os"
)

// NewLogger creates a new log.Logger which sends logs to stdout.
func NewLogger(prefix string) *log.Logger {
    return log.New(os.Stdout, prefix, 0)
}

