package spoke

import (
	"bitbucket.org/lewington/autoroller/globals"
)

// Spoke periodically pings an endpoint for
// health reports and plays alerts based on
// the current health status.
type Spoke interface {
	AddAlert(comp globals.Component, alert string)
	SpinUp()
	Close()
}
