package healthz

import "k8s.io/klog/v2"

var (
    currentStatus Status = StatusUp
    currentReason string = "service healthy"
)

// Status represents a state of either "up" or "down".
type Status string

const (
    // StatusUp represents the service being healthy.
    StatusUp Status = "up"

    // StatusDown represents the service being unhealthy.
    StatusDown Status = "down"
)

// SetStatus configures the status value (up/down). See [Status].
func SetStatus(status Status) {
    klog.Infof("health status changing: %s", status)
    currentStatus = status
}

// SetReason configures the reason for the current status.
func SetReason(reason string) {
    klog.Infof("health reason changing: %s", reason)
    currentReason = reason
}

// SetUp is a shortcut to run [SetStatus] and [SetReason] together.
func SetUp(reason string) {
    SetStatus(StatusUp)
    SetReason(reason)
}

// SetDown is a shortcut to run [SetStatus] and [SetReason] together.
func SetDown(reason string) {
    SetStatus(StatusDown)
    SetReason(reason)
}

