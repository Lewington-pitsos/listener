package update

import (
	"github.com/lewington/listener/globals"
	"github.com/lewington/listener/warn/health"
)

// Pinger sends system status updates from a given
// system to a given hub.
type Pinger interface {
	Update(component globals.Component, status health.Status) error
}
