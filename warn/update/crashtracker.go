package update

import (
	"time"

	"github.com/lewington/listener/warn/health"

	"github.com/lewington/listener/globals"
)

// CrashTracker keeps sending green updates to the
// hub. All it tracks is whether the system has crashed.
type CrashTracker struct {
	pinger    Pinger
	component globals.Component
	interval  time.Duration
}

// StartTracking sends a green status update for the
// tracker's component every interval.
func (t *CrashTracker) StartTracking() {
	for {
		t.pinger.Update(t.component, health.StatusGreen)
		time.Sleep(t.interval)
	}
}

// NewCrashTracker initializes a CrashTracker.
func NewCrashTracker(component globals.Component) *CrashTracker {
	return &CrashTracker{
		JSONHubPinger(),
		component,
		globals.HealthUpdateInterval,
	}
}
