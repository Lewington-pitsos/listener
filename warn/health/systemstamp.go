package health

import (
	"time"

	"github.com/lewington/listener/assist"
	"github.com/lewington/listener/globals"
)

// SystemStamp represents the state of a given system
// at a given timestamp.
type SystemStamp struct {
	System
	Timestamp time.Time
}

// NewSystemStamp initializes a SystemStamp stamped at the current time.
func NewSystemStamp(name globals.Component, status Status) *SystemStamp {
	return &SystemStamp{
		System{
			name,
			status,
		},
		assist.Timestamp(),
	}
}
