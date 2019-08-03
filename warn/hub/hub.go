package hub

// Hub keeps track of the overall system health
// and serves a representation of that health.
type Hub interface {
	Close()
}
