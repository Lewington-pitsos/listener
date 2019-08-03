package update

import (
	"bitbucket.org/lewington/autoroller/globals"
	"bitbucket.org/lewington/autoroller/warn/health"
)

// Pinger sends system status updates from a given
// system to a given hub.
type Pinger interface {
	Update(component globals.Component, status health.Status) error
}
