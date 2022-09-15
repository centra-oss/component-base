package logs

import (
	"flag"
	"fmt"
	"log"
	"time"

	"k8s.io/klog/v2"
)

var (
    packageFlags = flag.NewFlagSet("logging", flag.ContinueOnError)
    logFlushFreq time.Duration
)

func init() {
    klog.InitFlags(packageFlags)
	packageFlags.DurationVar(&logFlushFreq, "log-flush-frequency", 5 * time.Second, "Maximum number of seconds between log flushes")
}

// KlogWriter serves as a bridge between the standard log package and the glog package.
type KlogWriter struct{}

// Write implements the io.Writer interface.
func (writer KlogWriter) Write(data []byte) (n int, err error) {
	klog.InfoDepth(1, string(data))
	return len(data), nil
}

// NewLogger creates a new log.Logger which sends logs to stdout.
//
//  log := logs.NewLogger()
//  log.Println("Hello, world!")
func NewLogger(prefix string) *log.Logger {
    return log.New(KlogWriter{}, prefix, 0)
}

func InitLogs() {
    log.SetOutput(KlogWriter{})
    log.SetFlags(0)

    // Start flushing now.
    klog.StartFlushDaemon(logFlushFreq)

    klog.EnableContextualLogging(true)
}

func FlushLogs() {
    klog.Flush()
}

func GlogSetter(val string) (string, error) {
    var level klog.Level
    if err := level.Set(val); err != nil {
        return "", fmt.Errorf("failed to set klog.logging.verbosity %s: %v", val, err)
    }
    return fmt.Sprintf("successfully set klog.logging.verbosity to %s", val), nil
}

