package cli

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.h4n.io/centra/component-base/logs"
	"k8s.io/klog/v2"
)

// Run provides common boilerplate code around executing a cobra command.
func Run(cmd *cobra.Command) int { 
    if logsInitialized, err := run(cmd); err != nil {
        if !logsInitialized {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        } else {
            klog.ErrorS(err, "command failed")
        }
        return 1
    }

    return 0
}

func run(cmd *cobra.Command) (logsInitialized bool, err error) {
    rand.Seed(time.Now().UnixNano())
    defer logs.FlushLogs()

    // In all cases, errors are handled below.
    cmd.SilenceErrors = true

    switch {
    case cmd.PersistentPreRun != nil:
        pre := cmd.PersistentPreRun
        cmd.PersistentPreRun = func (cmd *cobra.Command, args []string) {
			logs.InitLogs()
			logsInitialized = true
            pre(cmd, args)
        }

    case cmd.PersistentPreRunE != nil:
        pre := cmd.PersistentPreRunE
        cmd.PersistentPreRunE = func (cmd *cobra.Command, args []string) error {
			logs.InitLogs()
			logsInitialized = true
            return pre(cmd, args)
        }

    default:
        cmd.PersistentPreRun = func (cmd *cobra.Command, args []string) {
			logs.InitLogs()
			logsInitialized = true
        }
    }

    err = cmd.Execute()

    return
}

