package spoke

import (
	"github.com/lewington/listener/globals"
)

// Spoke periodically pings an endpoint for
// health reports and plays alerts based on
// the current health status.
type Spoke interface {
	AddAlert(comp globals.Component, alert string)
	SpinUp()
	Close()
}
