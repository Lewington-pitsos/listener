package update

// Tracker is responsible for tracking the health of
// a system and relaying it to the hub.
type Tracker interface {
	StartTracking()
}
