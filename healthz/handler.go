package healthz

import (
	"encoding/json"
	"net/http"
)

// healthResponse represents the response that will be returned to the caller.
type healthResponse struct {
    Healthy bool
    Message string
}

// Handler is an HTTP handler function that presents the current status as
// configured by [SetStatus] and [SetReason].
func Handler(rw http.ResponseWriter, _ *http.Request) {
    text, err := json.Marshal(healthResponse{Healthy: currentStatus == StatusUp, Message: currentReason})
    if err != nil {
        rw.WriteHeader(http.StatusInternalServerError)
        rw.Write([]byte("Failed to marshal response object."))
    }

    // If the status is "down", set the HTTP response code accordingly.
    if currentStatus == StatusDown {
        rw.WriteHeader(http.StatusFailedDependency)
    }

    // Send the response JSON to the caller
    rw.Write(text)
}

