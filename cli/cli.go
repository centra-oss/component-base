package cli

import (
	"log"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// Run provides common boilerplate code around executing a cobra command.
func Run(cmd *cobra.Command) int { 
    if err := run(cmd); err != nil {
        log.Fatalf("command failed: %v", err)
        return 1
    }

    return 0
}

func run(cmd *cobra.Command) (err error) {
    rand.Seed(time.Now().UnixNano())

    // In all cases, errors are handled below.
    cmd.SilenceErrors = true

    switch {
    case cmd.PersistentPreRun != nil:
        pre := cmd.PersistentPreRun
        cmd.PersistentPreRun = func (cmd *cobra.Command, args []string) {
            pre(cmd, args)
        }

    case cmd.PersistentPreRunE != nil:
        pre := cmd.PersistentPreRunE
        cmd.PersistentPreRunE = func (cmd *cobra.Command, args []string) error {
            return pre(cmd, args)
        }

    default:
        cmd.PersistentPreRun = func (cmd *cobra.Command, args []string) {
        }
    }

    err = cmd.Execute()

    return
}

